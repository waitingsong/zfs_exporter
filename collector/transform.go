package collector

import (
	"fmt"
	"strconv"

	"github.com/waitingsong/zfs_exporter/v3/zfs"
)

type poolHealthCode int

const (
	poolOnline poolHealthCode = iota
	poolDegraded
	poolFaulted
	poolOffline
	poolUnavail
	poolRemoved
	poolSuspended
)

func transformNumeric(value string) (float64, error) {
	if value == `-` || value == `none` {
		return 0, nil
	}
	return strconv.ParseFloat(value, 64)
}

func transformHealthCode(status string) (float64, error) {
	var result poolHealthCode
	switch zfs.PoolStatus(status) {
	case zfs.PoolOnline:
		result = poolOnline
	case zfs.PoolDegraded:
		result = poolDegraded
	case zfs.PoolFaulted:
		result = poolFaulted
	case zfs.PoolOffline:
		result = poolOffline
	case zfs.PoolUnavail:
		result = poolUnavail
	case zfs.PoolRemoved:
		result = poolRemoved
	case zfs.PoolSuspended:
		result = poolSuspended
	default:
		return -1, fmt.Errorf(`unknown pool heath status: %s`, status)
	}

	return float64(result), nil
}

func transformBool(value string) (float64, error) {
	switch value {
	case `on`, `yes`, `enabled`, `active`:
		return 1, nil
	case `off`, `no`, `disabled`, `inactive`, `-`:
		return 0, nil
	}

	return -1, fmt.Errorf(`could not convert '%s' to bool`, value)
}

func transformPercentage(value string) (float64, error) {
	if len(value) > 0 && value[len(value)-1] == '%' {
		value = value[:len(value)-1]
	}
	v, err := transformNumeric(value)
	if err != nil {
		return -1, err
	}

	return v / 100, nil
}

func transformMultiplier(value string) (float64, error) {
	if len(value) > 0 && value[len(value)-1] == 'x' {
		value = value[:len(value)-1]
	}
	v, err := transformNumeric(value)
	if err != nil {
		return -1, err
	}
	return v, nil
}

type compressionCode int

// `The compression algorithm used for this dataset. [0: off, 1: on, 2: lz4, 3: zstd, 4: zstd-fast, 3xx: zstd-N, 4xxxx: zstd-fast-N].`,
const (
	compressOff compressionCode = iota
	compressOn
	compressLZ4
	compressZSTD
	compressZSTDFast
)

func transformCompression(algo string) (float64, error) {
	var result compressionCode
	switch zfs.CompressionAlgo(algo) {
	case zfs.CompressOff:
		result = compressOff

	case zfs.CompressOn:
		result = compressOn

	case zfs.CompressLZ4:
		result = compressLZ4

	case zfs.CompressZSTD:
		result = compressZSTD
	case zfs.CompressZSTDFast:
		result = compressZSTDFast

	default:
		// 4xxxx: zstd-fast-N, N=0-1000
		if len(algo) > 10 && algo[0:10] == `zstd-fast-` {
			if n, err := strconv.Atoi(algo[10:]); err == nil {
				return float64(40000 + n), nil
			}
		} else if len(algo) > 5 && algo[0:5] == `zstd-` { // 3xx: zstd-N, N=0-19
			if n, err := strconv.Atoi(algo[5:]); err == nil {
				return float64(300 + n), nil
			}
		}

		// return -1, fmt.Errorf(`unknown compress algo: %s`, algo)
		result = compressOn
	}

	return float64(result), nil
}

// `Handling of synchronous requests in this dataset. [1: latency, 2: throughput].`,
func transformLogbias(logbias string) (float64, error) {
	switch logbias {
	case `latency`:
		return 1, nil
	case `throughput`:
		return 2, nil
	}

	return -1, fmt.Errorf(`unknown logbias: %s`, logbias)
}
