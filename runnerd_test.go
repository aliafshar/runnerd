
package runnerd

import (
    . "launchpad.net/gocheck"
    "testing"
)

type RunnerSuite struct{}

func Test(t *testing.T) {
	Suite(&RunnerSuite{})
	TestingT(t)
}

func (s *RunnerSuite) TestHelloWorld(c *C) {
    c.Check(42, Equals, 42)
}
