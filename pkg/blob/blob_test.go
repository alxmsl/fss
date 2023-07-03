package blob

import (
	. "gopkg.in/check.v1"

	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type BlobSuite struct {
	cfg *BlobConfig
}

var _ = Suite(&BlobSuite{
	cfg: &BlobConfig{
		BucketName: "file:///tmp/test_crud_blob_model/subdir?create_dir=1",
		Timeout:    60 * time.Second,
	},
})

func (s *BlobSuite) SetUpSuite(c *C) {
	err := os.RemoveAll("/tmp/test_crud_blob_model")
	c.Assert(err, IsNil)
	err = os.MkdirAll("/tmp/test_crud_blob_model", 0755)
	c.Assert(err, IsNil)
	err = os.RemoveAll("/tmp/test_read_blob_contents")
	c.Assert(err, IsNil)
	err = os.MkdirAll("/tmp/test_read_blob_contents", 0755)
	c.Assert(err, IsNil)
}

func (s *BlobSuite) TearDownSuite(c *C) {
	err := os.RemoveAll("/tmp/test_crud_blob_model")
	c.Assert(err, IsNil)
	err = os.RemoveAll("/tmp/test_read_blob_contents")
	c.Assert(err, IsNil)
}

func (s *BlobSuite) TestConfig(c *C) {
	// test create
	var err = s.cfg.Put("test", strings.NewReader("test"))
	c.Assert(err, IsNil)
}

func (s *BlobSuite) TestReadBlobContents(c *C) {
	// test create
	err := s.cfg.Put("test", strings.NewReader("test"))
	c.Assert(err, IsNil)

	// test read
	buf := new(bytes.Buffer)
	err = s.cfg.Get("test", buf)

	c.Assert(err, IsNil)
	c.Assert(buf.String(), Equals, "test")
}
