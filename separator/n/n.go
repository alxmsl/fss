package n

import (
	"math"
)

type SeparatorN struct {
	n int64
}

func NewSeparatorN(n int64) *SeparatorN {
	return &SeparatorN{n}
}

func (s *SeparatorN) Separate(n int64) []int64 {
	var (
		part = float64(n) / float64(s.n)
		size = s.n
		i    int64
	)
	if part < 1. {
		size = n
		res := make([]int64, size)
		for i = 0; i < size; i += 1 {
			res[i] = 1
		}
		return res
	}

	var (
		res = make([]int64, size)

		ceil  = math.Ceil(part)
		floor = math.Floor(part)
		sum   int64
	)
	for i = 0; i < size; i += 1 {
		res[i] = int64(ceil)
		sum += res[i]
	}
	if sum == n {
		return res
	}
	for i = 0; i < size; i += 1 {
		res[size-i-1] = int64(floor)
		if sum-i-1 == n {
			break
		}
	}
	return res
}
