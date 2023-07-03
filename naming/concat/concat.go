package concat

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

const delimiter = "::"

type ConcatNaming struct {
	nonce string
	vv    []string
}

func FromName(name string) *ConcatNaming {
	var vv = strings.Split(name, delimiter)
	return &ConcatNaming{
		nonce: vv[0],
		vv:    vv[1:],
	}
}

func NewConcatNaming() *ConcatNaming {
	return &ConcatNaming{
		nonce: uuid.NewV4().String(),
		vv:    []string{},
	}
}

func (n *ConcatNaming) Add(v string) {
	n.vv = append(n.vv, v)
}

func (n *ConcatNaming) Name() string {
	return n.nonce
}

func (n *ConcatNaming) Parts() []string {
	return n.vv
}

func (n *ConcatNaming) String() string {
	var ss = append([]string{n.nonce}, n.vv...)
	return strings.Join(ss, delimiter)
}
