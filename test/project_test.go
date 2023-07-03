package test

import (
	. "gopkg.in/check.v1"

	"bytes"
	"io"
	"net/http"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type ProjectSuite struct{}

var _ = Suite(&ProjectSuite{})

func (s *ProjectSuite) TestService(c *C) {
	const data = `abcdefghijklmnopqrstuvwxyz0123456789`

	var res, err = http.Post("http://127.0.0.1:8080", "application/octet-stream", bytes.NewReader([]byte(data)))
	c.Assert(err, IsNil)
	c.Assert(res.StatusCode, Equals, http.StatusOK)
	bb, err := io.ReadAll(res.Body)
	c.Assert(err, IsNil)
	c.Assert(bb, Not(HasLen), 0)

	res, err = http.Get("http://127.0.0.1:8080/" + string(bb))
	c.Assert(err, IsNil)
	bb, err = io.ReadAll(res.Body)
	c.Assert(err, IsNil)

	c.Assert(string(bb), Equals, data)
}
