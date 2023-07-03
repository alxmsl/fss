package naming

import (
	"fmt"
)

type Interface interface {
	fmt.Stringer
	Add(string)
	Name() string
	Parts() []string
}
