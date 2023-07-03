package blob

import (
	_ "gocloud.dev/blob/fileblob"

	"context"
	"io"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"gocloud.dev/blob"
)

type BlobConfig struct {
	sync.RWMutex
	bucket *blob.Bucket

	BucketName string
	Timeout    time.Duration
}

func NewBlobConfig(bucketName string, timeout time.Duration) *BlobConfig {
	return &BlobConfig{
		BucketName: bucketName,
		Timeout:    timeout,
	}
}

func (cfg *BlobConfig) Context() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), cfg.Timeout)
	return ctx
}

func (cfg *BlobConfig) Put(name string, r io.Reader) error {
	b, err := cfg.Bucket()
	if err != nil {
		return err
	}
	w, err := b.NewWriter(cfg.Context(), name, &blob.WriterOptions{
		BeforeWrite: func(as func(interface{}) bool) error {
			var oh **storage.ObjectHandle
			switch {
			case as(&oh):
				*oh = (*oh).If(storage.Conditions{DoesNotExist: true})
			default:
				//@todo: for test purposes only
			}
			return nil
		},
	})
	if err != nil {
		return err
	}
	if _, err = io.Copy(w, r); err != nil {
		return err
	}
	return w.Close()
}

func (cfg *BlobConfig) Bucket() (*blob.Bucket, error) {
	cfg.RLock()
	if cfg.bucket != nil {
		defer cfg.RUnlock()
		return cfg.bucket, nil
	}
	cfg.RUnlock()

	cfg.Lock()
	defer cfg.Unlock()
	if cfg.bucket != nil {
		return cfg.bucket, nil
	}

	var err error

	cfg.bucket, err = blob.OpenBucket(cfg.Context(), cfg.BucketName)
	return cfg.bucket, err
}

func (cfg *BlobConfig) Get(name string, w io.Writer) error {
	b, err := cfg.Bucket()
	if err != nil {
		return err
	}

	r, err := b.NewReader(cfg.Context(), name, nil)
	if err != nil {
		return err
	}
	if _, err = io.Copy(w, r); err != nil {
		return err
	}
	return r.Close()
}
