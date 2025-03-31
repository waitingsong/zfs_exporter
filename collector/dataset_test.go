package collector

import (
	"context"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/waitingsong/zfs_exporter/v3/zfs"
	"github.com/waitingsong/zfs_exporter/v3/zfs/mock_zfs"
)

type datasetResults struct {
	name    string
	results map[string]string
}

func TestDatsetMetrics(t *testing.T) {
	testCases := []struct {
		name           string
		kinds          []zfs.DatasetKind
		pools          []string
		explicitPools  []string
		propsRequested []string
		metricNames    []string
		propsResults   map[string][]datasetResults
		metricResults  string
	}{
		{
			name:  `all metrics`,
			kinds: []zfs.DatasetKind{zfs.DatasetFilesystem},
			pools: []string{`testpool`},
			propsRequested: []string{
				`available`,
				`compression`,
				`compressratio`,
				`logbias`,
				`logicalused`,
				`logicalreferenced`,
				`primarycache`,
				`quota`,
				`refcompressratio`,
				`referenced`,
				`refquota`,
				`refreservation`,
				`reservation`,
				`snapshot_count`,
				`snapshot_limit`,
				`sync`,
				`used`,
				`usedbychildren`,
				`usedbydataset`,
				`usedbyrefreservation`,
				`usedbysnapshots`,
				`volsize`,
				`written`},
			metricNames: []string{
				`zfs_dataset_available_bytes`,
				`zfs_dataset_compression`,
				`zfs_dataset_compressratio`,
				`zfs_dataset_logbias`,
				`zfs_dataset_logical_used_bytes`,
				`zfs_dataset_logical_referenced_bytes`,
				`zfs_dataset_primarycache`,
				`zfs_dataset_quota_bytes`,
				`zfs_dataset_refcompressratio`,
				`zfs_dataset_referenced_bytes`,
				`zfs_dataset_referenced_quota_bytes`,
				`zfs_dataset_reservation_bytes`,
				`zfs_dataset_snapshot_count_total`,
				`zfs_datset_snapshot_limit_total`,
				`zfs_dataset_sync`,
				`zfs_dataset_used_bytes`,
				`zfs_dataset_used_by_children_bytes`,
				`zfs_dataset_used_by_datset_bytes`,
				`zfs_datset_used_by_referenced_reservation_bytes`,
				`zfs_dataset_used_by_snapshot_bytes`,
				`zfs_dataset_volume_size_bytes`,
				`zfs_dataset_written_bytes`},
			propsResults: map[string][]datasetResults{
				`testpool`: {
					{
						name: `testpool/test`,
						results: map[string]string{
							`available`:            `1024`,
							`compression`:          `zstd-2`,
							`compressratio`:        `2.50`,
							`logbias`:              `latency`,
							`logicalused`:          `1024`,
							`logicalreferenced`:    `512`,
							`primarycache`:         `all`,
							`quota`:                `512`,
							`refcompressratio`:     `24.00`,
							`referenced`:           `1024`,
							`refreservation`:       `1024`,
							`reservation`:          `1024`,
							`snapshot_count`:       `12`,
							`snapshot_limit`:       `24`,
							`sync`:                 `standard`,
							`used`:                 `1024`,
							`usedbychildren`:       `1024`,
							`usedbydataset`:        `1024`,
							`usedbyrefreservation`: `1024`,
							`usedbysnapshots`:      `1024`,
							`volsize`:              `1024`,
							`written`:              `1024`,
						},
					},
				},
			},
			metricResults: `# HELP zfs_dataset_available_bytes The amount of space in bytes available to the dataset and all its children.
# TYPE zfs_dataset_available_bytes gauge
zfs_dataset_available_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
# HELP zfs_dataset_compressratio For non-snapshots, the compression ratio achieved for the used space of this dataset, For snapshots, the compressratio is the same as the refcompressratio property.
# TYPE zfs_dataset_compressratio gauge
zfs_dataset_compressratio{name="testpool/test",pool="testpool",type="filesystem"} 2.5
# HELP zfs_dataset_logical_used_bytes The amount of space in bytes that is "logically" consumed by this dataset and all its descendents. See the "used_bytes" property.
# TYPE zfs_dataset_logical_used_bytes gauge
zfs_dataset_logical_used_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
# HELP zfs_dataset_logical_referenced_bytes The amount of space that is "logically" accessible by this dataset. See the "referenced_bytes" property.
# TYPE zfs_dataset_logical_referenced_bytes gauge
zfs_dataset_logical_referenced_bytes{name="testpool/test",pool="testpool",type="filesystem"} 512
# HELP zfs_dataset_quota_bytes The maximum amount of space in bytes this dataset and its descendents can consume.
# TYPE zfs_dataset_quota_bytes gauge
zfs_dataset_quota_bytes{name="testpool/test",pool="testpool",type="filesystem"} 512
# HELP zfs_dataset_referenced_bytes The amount of data in bytes that is accessible by this dataset, which may or may not be shared with other datasets in the pool.
# TYPE zfs_dataset_referenced_bytes gauge
zfs_dataset_referenced_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
# HELP zfs_dataset_refcompressratio The compression ratio achieved for the referenced space of this dataset, expressed as a multiplier.
# TYPE zfs_dataset_refcompressratio gauge
zfs_dataset_refcompressratio{name="testpool/test",pool="testpool",type="filesystem"} 24
# HELP zfs_dataset_reservation_bytes The minimum amount of space in bytes guaranteed to a dataset and its descendants.
# TYPE zfs_dataset_reservation_bytes gauge
zfs_dataset_reservation_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
# HELP zfs_dataset_snapshot_count_total The total number of snapshots that exist under this location in the dataset tree. This value is only available when a snapshot_limit has been set somewhere in the tree under which the dataset resides.
# TYPE zfs_dataset_snapshot_count_total gauge
zfs_dataset_snapshot_count_total{name="testpool/test",pool="testpool",type="filesystem"} 12
# HELP zfs_dataset_used_by_children_bytes The amount of space in bytes used by children of this dataset, which would be freed if all the dataset's children were destroyed.
# TYPE zfs_dataset_used_by_children_bytes gauge
zfs_dataset_used_by_children_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
# HELP zfs_dataset_used_by_snapshot_bytes The amount of space in bytes consumed by snapshots of this dataset.
# TYPE zfs_dataset_used_by_snapshot_bytes gauge
zfs_dataset_used_by_snapshot_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
# HELP zfs_dataset_used_bytes The amount of space in bytes consumed by this dataset and all its descendents.
# TYPE zfs_dataset_used_bytes gauge
zfs_dataset_used_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
# HELP zfs_dataset_volume_size_bytes The logical size in bytes of this volume.
# TYPE zfs_dataset_volume_size_bytes gauge
zfs_dataset_volume_size_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
# HELP zfs_dataset_written_bytes The amount of referenced space in bytes written to this dataset since the previous snapshot.
# TYPE zfs_dataset_written_bytes gauge
zfs_dataset_written_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
# HELP zfs_dataset_compression The compression algorithm used for this dataset. [0: off, 1: on, 2: lz4, 3: zstd, 4: zstd-fast, 3xx: zstd-N, 4xxxx: zstd-fast-N].
# TYPE zfs_dataset_compression gauge
zfs_dataset_compression{name="testpool/test",pool="testpool",type="filesystem"} 302
# HELP zfs_dataset_logbias Handling of synchronous requests in this dataset. [1: latency, 2: throughput].
# TYPE zfs_dataset_logbias gauge
zfs_dataset_logbias{name="testpool/test",pool="testpool",type="filesystem"} 1
# HELP zfs_dataset_sync The sync behavior of this dataset [1: standard, 2: always, 0: disabled].
# TYPE zfs_dataset_sync gauge
zfs_dataset_sync{name="testpool/test",pool="testpool",type="filesystem"} 1
# HELP zfs_dataset_primarycache What is cached in the primary cache (ARC) [1: all, 2: metadata, 0: none].
# TYPE zfs_dataset_primarycache gauge
zfs_dataset_primarycache{name="testpool/test",pool="testpool",type="filesystem"} 1
`,
		},
		{
			name:           `multiple pools`,
			kinds:          []zfs.DatasetKind{zfs.DatasetFilesystem},
			pools:          []string{`testpool1`, `testpool2`},
			propsRequested: []string{`available`},
			metricNames:    []string{`zfs_dataset_available_bytes`},
			propsResults: map[string][]datasetResults{
				`testpool1`: {
					{
						name: `testpool1/test`,
						results: map[string]string{
							`available`: `1024`,
						},
					},
				},
				`testpool2`: {
					{
						name: `testpool2/test`,
						results: map[string]string{
							`available`: `1024`,
						},
					},
				},
			},
			metricResults: `# HELP zfs_dataset_available_bytes The amount of space in bytes available to the dataset and all its children.
# TYPE zfs_dataset_available_bytes gauge
zfs_dataset_available_bytes{name="testpool1/test",pool="testpool1",type="filesystem"} 1024
zfs_dataset_available_bytes{name="testpool2/test",pool="testpool2",type="filesystem"} 1024
`,
		},
		{
			name:           `explicit pools`,
			kinds:          []zfs.DatasetKind{zfs.DatasetFilesystem},
			pools:          []string{`testpool1`, `testpool2`},
			explicitPools:  []string{`testpool1`},
			propsRequested: []string{`available`},
			metricNames:    []string{`zfs_dataset_available_bytes`},
			propsResults: map[string][]datasetResults{
				`testpool1`: {
					{
						name: `testpool1/test`,
						results: map[string]string{
							`available`: `1024`,
						},
					},
				},
				`testpool2`: {
					{
						name: `testpool2/test`,
						results: map[string]string{
							`available`: `1024`,
						},
					},
				},
			},
			metricResults: `# HELP zfs_dataset_available_bytes The amount of space in bytes available to the dataset and all its children.
# TYPE zfs_dataset_available_bytes gauge
zfs_dataset_available_bytes{name="testpool1/test",pool="testpool1",type="filesystem"} 1024
`,
		},
		{
			name:           `multiple collectors`,
			kinds:          []zfs.DatasetKind{zfs.DatasetFilesystem, zfs.DatasetSnapshot, zfs.DatasetVolume},
			pools:          []string{`testpool`},
			propsRequested: []string{`available`},
			metricNames:    []string{`zfs_dataset_available_bytes`},
			propsResults: map[string][]datasetResults{
				`testpool`: {
					{
						name: `testpool/test`,
						results: map[string]string{
							`available`: `1024`,
						},
					},
				},
			},
			metricResults: `# HELP zfs_dataset_available_bytes The amount of space in bytes available to the dataset and all its children.
# TYPE zfs_dataset_available_bytes gauge
zfs_dataset_available_bytes{name="testpool/test",pool="testpool",type="filesystem"} 1024
zfs_dataset_available_bytes{name="testpool/test",pool="testpool",type="snapshot"} 1024
zfs_dataset_available_bytes{name="testpool/test",pool="testpool",type="volume"} 1024
`,
		},
		{
			name:           `unsupported metric`,
			kinds:          []zfs.DatasetKind{zfs.DatasetFilesystem},
			pools:          []string{`testpool`},
			propsRequested: []string{`unsupported`},
			metricNames:    []string{`zfs_dataset_unsupported`},
			propsResults: map[string][]datasetResults{
				`testpool`: {
					{
						name: `testpool/test`,
						results: map[string]string{
							`unsupported`: `1024`,
						},
					},
				},
			},
			metricResults: `# HELP zfs_dataset_unsupported !!! This property is unsupported, results are likely to be undesirable, please file an issue at https://github.com/waitingsong/zfs_exporter/issues to have this property supported !!!
# TYPE zfs_dataset_unsupported gauge
zfs_dataset_unsupported{name="testpool/test",pool="testpool",type="filesystem"} 1024
`,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctrl, ctx := gomock.WithContext(context.Background(), t)
			zfsClient := mock_zfs.NewMockClient(ctrl)
			config := defaultConfig(zfsClient)
			if tc.explicitPools != nil {
				config.Pools = tc.explicitPools
			}

			zfsClient.EXPECT().PoolNames().Return(tc.pools, nil).Times(1)
			collector, err := NewZFS(config)
			if err != nil {
				t.Fatal(err)
			}
			collector.Collectors = make(map[string]State)

			for _, kind := range tc.kinds {
				switch kind {
				case zfs.DatasetFilesystem:
					collector.Collectors[`dataset-filesystem`] = State{
						Name:       "dataset-filesystem",
						Enabled:    boolPointer(true),
						Properties: stringPointer(strings.Join(tc.propsRequested, `,`)),
						factory:    newFilesystemCollector,
					}
				case zfs.DatasetSnapshot:
					collector.Collectors[`dataset-snapshot`] = State{
						Name:       "dataset-snapshot",
						Enabled:    boolPointer(true),
						Properties: stringPointer(strings.Join(tc.propsRequested, `,`)),
						factory:    newSnapshotCollector,
					}
				case zfs.DatasetVolume:
					collector.Collectors[`dataset-volume`] = State{
						Name:       "dataset-volume",
						Enabled:    boolPointer(true),
						Properties: stringPointer(strings.Join(tc.propsRequested, `,`)),
						factory:    newVolumeCollector,
					}
				}
				for _, pool := range tc.pools {
					if tc.explicitPools != nil {
						wanted := false
						for _, explicit := range tc.explicitPools {
							if pool == explicit {
								wanted = true
							}
							break
						}
						if !wanted {
							continue
						}
					}
					zfsDatasetResults := make([]zfs.DatasetProperties, len(tc.propsResults[pool]))
					for i, propResults := range tc.propsResults[pool] {
						zfsDatasetProperties := mock_zfs.NewMockDatasetProperties(ctrl)
						zfsDatasetProperties.EXPECT().DatasetName().Return(propResults.name).Times(2)
						zfsDatasetProperties.EXPECT().Properties().Return(propResults.results).Times(1)
						zfsDatasetResults[i] = zfsDatasetProperties
					}
					zfsDatasets := mock_zfs.NewMockDatasets(ctrl)
					zfsDatasets.EXPECT().Properties(tc.propsRequested).Return(zfsDatasetResults, nil).Times(1)
					zfsClient.EXPECT().Datasets(pool, kind).Return(zfsDatasets).Times(1)
				}
			}

			if err = callCollector(ctx, collector, []byte(tc.metricResults), tc.metricNames); err != nil {
				t.Fatal(err)
			}
		})
	}
}
