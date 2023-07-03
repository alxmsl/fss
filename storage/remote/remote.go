package remote

import (
	"errors"
	"io"
	"math"
	"net/http"
	"strconv"
)

type Storage struct {
	name, addr string
}

func NewStorage(name, addr string) *Storage {
	return &Storage{
		name: name,
		addr: addr,
	}
}

func (s *Storage) Get(name string, w io.Writer) error {
	var res, err = http.Get(s.addr + "/GetObject/" + name)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New("wrong status code")
	}
	_, err = io.Copy(w, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Name() string {
	return s.name
}

func (s *Storage) Put(name string, r io.Reader) error {
	var res, err = http.Post(s.addr+"/PutObject/"+name, "application/octet-stream", r)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("wrong status code")
	}
	return nil
}

func (s *Storage) Size() int64 {
	var res, err = http.Get(s.addr + "/Size")
	if err != nil {
		return math.MaxInt64
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return math.MaxInt64
	}
	bb, err := io.ReadAll(res.Body)
	if err != nil {
		return math.MaxInt64
	}
	size, err := strconv.Atoi(string(bb))
	if err != nil {
		return math.MaxInt64
	}
	return int64(size)
}
