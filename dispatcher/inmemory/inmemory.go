package inmemory

import (
	"bytes"
	"io"
	"math"
	"sync"

	"github.com/alxmsl/fss/naming"
	"github.com/alxmsl/fss/storage"
)

type Dispatcher struct {
	sync.Mutex
	sm map[string]*state
	ss []*state
}

func NewDispatcher(storages []storage.Interface) *Dispatcher {
	var d = &Dispatcher{
		sm: map[string]*state{},
		ss: make([]*state, 0, len(storages)),
	}
	for _, s := range storages {
		var ss = &state{storage: s}
		if _, ok := d.sm[s.Name()]; ok {
			panic("duplicated storage")
		}
		d.sm[s.Name()] = ss
		d.ss = append(d.ss, ss)
	}
	return d
}

func (d *Dispatcher) DownloadPart(name naming.Interface, part string, w io.Writer) error {
	if _, ok := d.sm[part]; !ok {
		panic("storage not found")
	}
	return d.sm[part].storage.Get(name.Name(), w)
}

func (d *Dispatcher) UploadPart(name naming.Interface, bb []byte) error {
	var (
		size = int64(len(bb))
		idx  = d.reserve(size)
		err  = d.ss[idx].storage.Put(name.Name(), bytes.NewReader(bb))
	)
	defer d.release(idx, size)
	if err != nil {
		return err
	}
	name.Add(d.ss[idx].storage.Name())
	return nil
}

func (d *Dispatcher) reserve(size int64) int {
	d.Lock()
	defer d.Unlock()

	var (
		min = int64(math.MaxInt64)
		idx int
	)
	for i, s := range d.ss {
		if s.Size() < min {
			min = s.Size()
			idx = i
		}
	}
	d.ss[idx].reserve += size
	return idx
}

func (d *Dispatcher) release(idx int, size int64) {
	d.Lock()
	defer d.Unlock()
	if d.ss[idx].reserve < size {
		panic("insufficient reserve")
	}
	d.ss[idx].reserve -= size
}
