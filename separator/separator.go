package separator

import "github.com/alxmsl/fss/separator/n"

type Interface interface {
	Separate(int64) []int64
}

var DefaultSeparator = n.NewSeparatorN(6)
