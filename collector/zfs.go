package collector

import (
	"context"
	"log/slog"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/waitingsong/zfs_exporter/v3/zfs"
)

type regexpCollection []*regexp.Regexp

func (c regexpCollection) MatchString(input string) bool {
	for _, r := range c {
		if r.MatchString(input) {
			return true
		}
	}

	return false
}

// ZFSConfig configures a ZFS collector
type ZFSConfig struct {
	DisableMetrics bool
	Deadline       time.Duration
	Pools          []string
	Excludes       []string
	Logger         *slog.Logger
	ZFSClient      zfs.Client
}

// ZFS collector
type ZFS struct {
	Pools          []string
	Collectors     map[string]State
	client         zfs.Client
	disableMetrics bool
	deadline       time.Duration
	cache          *metricCache
	ready          chan struct{}
	logger         *slog.Logger
	excludes       regexpCollection
}

// Describe implements the prometheus.Collector interface.
func (c *ZFS) Describe(ch chan<- *prometheus.Desc) {
	if !c.disableMetrics {
		ch <- scrapeDurationDesc
		ch <- scrapeSuccessDesc
	}

	for _, state := range c.Collectors {
		if !*state.Enabled {
			continue
		}

		collector, err := state.factory(c.logger, c.client, strings.Split(*state.Properties, `,`))
		if err != nil {
			continue
		}
		collector.describe(ch)
	}
}

// Collect implements the prometheus.Collector interface.
func (c *ZFS) Collect(ch chan<- prometheus.Metric) {
	select {
	case <-c.ready:
	default:
		c.sendCached(ch, make(map[string]struct{}))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), c.deadline)
	defer cancel()

	cache := newMetricCache()
	proxy := make(chan metric)
	// Synchronize on collector completion.
	wg := sync.WaitGroup{}
	wg.Add(len(c.Collectors))
	// Synchonize after timeout event, ensuring no writers are still active when we return control.
	timeout := make(chan struct{})
	finalized := make(chan struct{})
	finalize := func() {
		select {
		case <-finalized:
		default:
			close(finalized)
		}
	}

	// Close the proxy channel upon collector completion.
	go func() {
		wg.Wait()
		close(proxy)
	}()

	// Cache metrics as they come in via the proxy channel, and ship them out if we've not exceeded the deadline.
	go func() {
		for metric := range proxy {
			cache.add(metric)
			select {
			case <-timeout:
				finalize()
			default:
				ch <- metric.prometheus
			}
		}
		// Signal completion and update full cache.
		c.cache.replace(cache)
		cancel()
		// Notify next collection that we're ready to collect again
		c.ready <- struct{}{}
	}()

	pools, poolErr := c.getPools(c.Pools)

	for name, state := range c.Collectors {
		if !*state.Enabled {
			wg.Done()
			continue
		}

		if poolErr != nil {
			c.publishCollectorMetrics(ctx, name, poolErr, 0, proxy)
			wg.Done()
			continue
		}

		collector, err := state.factory(c.logger, c.client, strings.Split(*state.Properties, `,`))
		if err != nil {
			c.logger.Error("Error instantiating collector", "collector", name, "err", err)
			wg.Done()
			continue
		}
		go func(name string, collector Collector) {
			c.execute(ctx, name, collector, proxy, pools)
			wg.Done()
		}(name, collector)
	}

	// Wait for completion or timeout
	<-ctx.Done()
	err := ctx.Err()
	if err == context.Canceled {
		finalize()
	} else if err != nil {
		// Upon exceeding deadline, send cached data for any metrics that have not already been reported.
		close(timeout) // assert timeout for flow control in other goroutines
		c.cache.merge(cache)
		cacheIndex := cache.index()
		c.sendCached(ch, cacheIndex)
	}
	// Ensure there are no in-flight writes to the upstream channel
	<-finalized
}

// sendCached values that do not appear in the current cacheIndex.
func (c *ZFS) sendCached(ch chan<- prometheus.Metric, cacheIndex map[string]struct{}) {
	c.cache.RLock()
	defer c.cache.RUnlock()
	for name, metric := range c.cache.cache {
		if _, ok := cacheIndex[name]; ok {
			continue
		}
		ch <- metric
	}
}

func (c *ZFS) getPools(pools []string) ([]string, error) {
	poolNames, err := c.client.PoolNames()
	if err != nil {
		return nil, err
	}
	// Return all pools if not explicitly configured.
	if len(pools) == 0 {
		return poolNames, nil
	}

	// Configured pools may not exist, so append available pools as they're found, rather than allocating up front.
	result := make([]string, 0)
	for _, want := range pools {
		found := false
		for _, avail := range poolNames {
			if want == avail {
				result = append(result, want)
				found = true
				break
			}
		}
		if !found {
			c.logger.Warn("Pool unavailable", "pool", want)
		}
	}

	return result, nil
}

func (c *ZFS) execute(ctx context.Context, name string, collector Collector, ch chan<- metric, pools []string) {
	begin := time.Now()
	err := collector.update(ch, pools, c.excludes)
	duration := time.Since(begin)

	c.publishCollectorMetrics(ctx, name, err, duration, ch)
}

func (c *ZFS) publishCollectorMetrics(ctx context.Context, name string, err error, duration time.Duration, ch chan<- metric) {
	var success float64

	if err != nil {
		c.logger.Error("Executing collector", "status", "error", "collector", name, "durationSeconds", duration.Seconds(), "err", err)
		success = 0
	} else {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
			err = nil
		}
		if err != nil && err != context.Canceled {
			c.logger.Warn("Executing collector", "status", "delayed", "collector", name, "durationSeconds", duration.Seconds(), "err", ctx.Err())
			success = 0
		} else {
			c.logger.Debug("Executing collector", "status", "ok", "collector", name, "durationSeconds", duration.Seconds())
			success = 1
		}
	}

	if c.disableMetrics {
		return
	}
	ch <- metric{
		name:       scrapeDurationDescName,
		prometheus: prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, duration.Seconds(), name),
	}
	ch <- metric{
		name:       scrapeSuccessDescName,
		prometheus: prometheus.MustNewConstMetric(scrapeSuccessDesc, prometheus.GaugeValue, success, name),
	}
}

// NewZFS instantiates a ZFS collector with the provided ZFSConfig
func NewZFS(config ZFSConfig) (*ZFS, error) {
	sort.Strings(config.Pools)
	sort.Strings(config.Excludes)
	excludes := make(regexpCollection, len(config.Excludes))
	for i, v := range config.Excludes {
		excludes[i] = regexp.MustCompile(v)
	}
	ready := make(chan struct{}, 1)
	ready <- struct{}{}
	return &ZFS{
		disableMetrics: config.DisableMetrics,
		client:         config.ZFSClient,
		deadline:       config.Deadline,
		Pools:          config.Pools,
		Collectors:     collectorStates,
		excludes:       excludes,
		cache:          newMetricCache(),
		ready:          ready,
		logger:         config.Logger,
	}, nil
}
