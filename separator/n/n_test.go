package n

import (
	. "gopkg.in/check.v1"

	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type SeparatorSuite struct {
	sep *SeparatorN
}

var _ = Suite(&SeparatorSuite{
	sep: NewSeparatorN(6),
})

func (s *SeparatorSuite) TestSeparate(c *C) {
	var ds = []struct {
		v int64
		r []int64
	}{
		{0, []int64{}},
		{1, []int64{1}},
		{2, []int64{1, 1}},
		{3, []int64{1, 1, 1}},
		{4, []int64{1, 1, 1, 1}},
		{5, []int64{1, 1, 1, 1, 1}},
		{6, []int64{1, 1, 1, 1, 1, 1}},
		{7, []int64{2, 1, 1, 1, 1, 1}},
		{8, []int64{2, 2, 1, 1, 1, 1}},
		{9, []int64{2, 2, 2, 1, 1, 1}},
		{10, []int64{2, 2, 2, 2, 1, 1}},
		{11, []int64{2, 2, 2, 2, 2, 1}},
		{12, []int64{2, 2, 2, 2, 2, 2}},
		{13, []int64{3, 2, 2, 2, 2, 2}},
	}
	for _, d := range ds {
		var pp = s.sep.Separate(d.v)
		c.Assert(pp, DeepEquals, d.r)
	}
}
