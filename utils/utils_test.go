package utils

import (
	"testing"

	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
)

func TestCleanPath(t *testing.T) {
	a := assertions.New(t)

	a.So(CleanPaths("a", "b", "c", "d.e", "a", "b.c"), should.Resemble, []string{"a", "b", "c", "d.e"})
	a.So(PathsWithPrefix("sub", "a", "b.c"), should.Resemble, []string{"sub.a", "sub.b.c"})
	a.So(PathsWithoutPrefix("sub", "sub.a", "sub.b.c", "d.e"), should.Resemble, []string{"a", "b.c"})
}
