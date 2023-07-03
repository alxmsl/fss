package local

import (
	"io"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/alxmsl/fss/pkg/blob"
)

type Storage struct {
	cfg  *blob.BlobConfig
	dir  string
	name string
}

func NewStorage(name, dir string, timeout time.Duration) *Storage {
	return &Storage{
		cfg:  blob.NewBlobConfig("file://"+dir+"?create_dir=1", timeout),
		dir:  dir,
		name: name,
	}
}

func (s *Storage) Get(name string, w io.Writer) error {
	return s.cfg.Get(name, w)
}

func (s *Storage) Name() string {
	return s.name
}

func (s *Storage) Put(name string, r io.Reader) error {
	return s.cfg.Put(name, r)
}

func (s *Storage) Size() int64 {
	var size int64
	err := filepath.Walk(s.dir, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		return math.MaxInt64
	}
	return size
}
