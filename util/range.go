package util

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/rs/zerolog/log"
)

type AddMode string

const (
	MinMax      AddMode = "min-max"
	Consecutive AddMode = "consecutive"
)

type Range struct {
	From, To int
}

func (r Range) String() string {
	return fmt.Sprintf("%d-%d", r.From, r.To)
}

func (r Range) StringValues(format string) (string, string) {
	return fmt.Sprintf(format, r.From), fmt.Sprintf(format, r.To)
}

func (r Range) Contains(x int) bool {
	return x >= r.From && x <= r.To
}

func (r Range) Add(x int, mode AddMode, warnIfOverlaps bool) (Range, bool, error) {
	const semLogContext = "opem-util::range-add"

	if r.Contains(x) {
		if warnIfOverlaps {
			err := errors.New("range overlap")
			log.Warn().Err(err).Interface("range", r).Int("value", x).Msg(semLogContext)
		}
		return r, true, nil
	}

	if x < r.From {
		switch mode {
		case MinMax:
			return Range{x, r.To}, true, nil
		case Consecutive:
			if x == (r.From - 1) {
				return Range{x, r.To}, true, nil
			}
		default:
			err := errors.New("invalid range maode")
			log.Error().Err(err).Msg(semLogContext)
			return r, false, err
		}
	}

	if x > r.To {
		switch mode {
		case MinMax:
			return Range{r.From, x}, true, nil
		case Consecutive:
			if x == (r.To + 1) {
				return Range{r.From, x}, true, nil
			}
		default:
			err := errors.New("invalid range maode")
			log.Error().Err(err).Msg(semLogContext)
			return r, false, err
		}
	}

	return Range{From: x, To: x}, false, nil
}

type RangeSet struct {
	Ranges []Range
}

func (rs *RangeSet) AddNumericString(x string, mode AddMode, warnIfOverlaps bool) error {
	i, err := strconv.Atoi(x)
	if err != nil {
		return err
	}

	return rs.Add(i, mode, warnIfOverlaps)
}

func (rs *RangeSet) Add(x int, mode AddMode, warnIfOverlaps bool) error {
	const semLogContext = "opem-util::range-set-add"

	if len(rs.Ranges) == 0 {
		rs.Ranges = append(rs.Ranges, Range{x, x})
	}

	for ri, r := range rs.Ranges {
		nr, ok, err := r.Add(x, mode, warnIfOverlaps)
		if err != nil {
			return err
		}

		if ok {
			rs.Ranges[ri] = nr
			return nil
		}

	}

	rs.Ranges = append(rs.Ranges, Range{x, x})
	return nil
}

func (rs *RangeSet) Defragment() []Range {
	if len(rs.Ranges) < 2 {
		return rs.Ranges
	}

	sort.Slice(rs.Ranges, func(i, j int) bool {
		return rs.Ranges[i].From < rs.Ranges[j].From
	})

	var res []Range
	rarr := rs.Ranges
	for i := 0; i < len(rarr); {
		nr := rarr[i]
		for j := i + 1; j <= len(rarr)-1; j++ {
			if rs.Ranges[j].From == nr.To+1 {
				nr = Range{nr.From, rs.Ranges[j].To}
				i = i + 1
			} else {
				break
			}
		}

		i = i + 1
		res = append(res, nr)
	}

	return res
}
