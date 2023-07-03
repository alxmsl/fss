package storage

import (
	"io"
)

type Interface interface {
	Get(name string, w io.Writer) error
	Name() string
	Put(name string, r io.Reader) error
	Size() int64
}
