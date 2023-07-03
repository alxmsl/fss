package local

import (
	. "gopkg.in/check.v1"
	"math"

	"bytes"
	"os"
	"strings"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type BlobSuite struct{}

var _ = Suite(&BlobSuite{})

func (s *BlobSuite) TestLocalStorage(c *C) {
	var err = os.RemoveAll("/tmp/local_storage_1")
	c.Assert(err, IsNil)

	var ls = NewStorage("local_storage_1", "/tmp/local_storage_1", time.Second)
	c.Assert(ls.Name(), Equals, "local_storage_1")
	c.Assert(ls.Size(), Equals, int64(math.MaxInt64))

	err = ls.Put("file_1", strings.NewReader("test_1"))
	c.Assert(err, IsNil)
	c.Assert(ls.Size(), Equals, int64(6+212))

	err = ls.Put("file_2", strings.NewReader("test_10"))
	c.Assert(err, IsNil)
	c.Assert(ls.Size(), Equals, int64(6+212+7+212))

	var buf = new(bytes.Buffer)
	err = ls.Get("file_2", buf)
	c.Assert(err, IsNil)
	c.Assert(buf.String(), Equals, "test_10")

	buf = new(bytes.Buffer)
	err = ls.Get("file_1", buf)
	c.Assert(err, IsNil)
	c.Assert(buf.String(), Equals, "test_1")
}
