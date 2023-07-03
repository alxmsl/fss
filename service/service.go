package service

import (
	"errors"
	"io"
	"log"

	"github.com/alxmsl/fss/dispatcher"
	"github.com/alxmsl/fss/naming/concat"
	"github.com/alxmsl/fss/separator"
)

type Interface interface {
	Download(string) (io.Writer, error)
	Upload(io.Reader, int64) (string, error)
}

type Service struct {
	dispatcher dispatcher.Interface
	separator  separator.Interface
}

func New(dispatcher dispatcher.Interface, separator separator.Interface) *Service {
	return &Service{
		dispatcher: dispatcher,
		separator:  separator,
	}
}

func (s *Service) Upload(src io.Reader, size int64) (string, error) {
	var (
		parts  = s.separator.Separate(size)
		naming = concat.NewConcatNaming()
	)
	for i := 0; i < len(parts); i += 1 {
		var (
			partSize = parts[i]
			bb       = make([]byte, partSize)
			n, err   = src.Read(bb)
		)
		if err != nil && !errors.Is(err, io.EOF) {
			return "", err
		}
		if int64(n) < parts[i] {
			return "", err
		}

		err = s.dispatcher.UploadPart(naming, bb)
		if err != nil {
			return "", err
		}
	}
	return naming.String(), nil
}

func (s *Service) Download(name string, dst io.Writer) error {
	var (
		naming = concat.FromName(name)
		parts  = naming.Parts()
		err    error
	)
	for _, part := range parts {
		log.Println(part)
		err = s.dispatcher.DownloadPart(naming, part, dst)
		if err != nil {
			return err
		}
	}
	return nil
}
