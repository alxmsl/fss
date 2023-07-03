package inmemory

import "github.com/alxmsl/fss/storage"

type state struct {
	storage storage.Interface
	reserve int64
}

func (s *state) Size() int64 {
	return s.storage.Size() + s.reserve
}
