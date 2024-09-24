package g_test

import (
	"errors"
	"testing"

	"github.com/avila-r/g"
)

func TestIf(t *testing.T) {
	{
		i1, i2 := 1, 2
		exp, got := i1, g.If(true, i1, i2)
		if got != exp {
			t.Errorf("[int] Expected %d, got: %d", exp, got)
		}
		exp, got = i2, g.If(false, i1, i2)
		if got != exp {
			t.Errorf("[int] Expected %d, got: %d", exp, got)
		}
	}

	{
		s1, s2 := "first", "second"
		exp, got := s1, g.If(true, s1, s2)
		if got != exp {
			t.Errorf("[string] Expected %s, got: %s", exp, got)
		}
		exp, got = s2, g.If(false, s1, s2)
		if got != exp {
			t.Errorf("[string] Expected %s, got: %s", exp, got)
		}
	}
}

func TestPtr(t *testing.T) {
	s := "a"
	sp := g.Ptr(s)
	if *sp != s {
		t.Errorf("Ptr[string] failed")
	}

	i := 2
	ip := g.Ptr(i)
	if *ip != i {
		t.Errorf("Ptr[int] failed")
	}
}

func TestMust(t *testing.T) {
	i := 1
	if got := g.Must(i, nil); got != i {
		t.Errorf("Must[int] failed")
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic")
			}
		}()
		g.Must(i, errors.New("test")) // Expecting panic
		t.Error("Not expected to reach this")
	}()
}

func manyResults() (i, j, k int, s string, f float64) {
	return 1, 2, 3, "four", 5.0
}

func TestFirst(t *testing.T) {
	exp, got := 1, g.First(manyResults())
	if got != exp {
		t.Errorf("Expected %d, got: %d", exp, got)
	}
}

func TestSecond(t *testing.T) {
	exp, got := 2, g.Second(manyResults())
	if got != exp {
		t.Errorf("Expected %d, got: %d", exp, got)
	}
}
func TestThird(t *testing.T) {
	exp, got := 3, g.Third(manyResults())
	if got != exp {
		t.Errorf("Expected %d, got: %d", exp, got)
	}
}

func TestCoalesce(t *testing.T) {
	p1, p2 := g.Ptr(1), g.Ptr(2)

	cases := []struct {
		name     string
		exp, got any
	}{
		{
			"strings",
			"1", g.Coalesce("", "1", "2"),
		},
		{
			"strings first",
			"1", g.Coalesce("1", "2", "3"),
		},
		{
			"strings last",
			"1", g.Coalesce("", "", "1"),
		},
		{
			"strings all zero",
			"", g.Coalesce("", "", ""),
		},
		{
			"strings no args",
			"", g.Coalesce[string](),
		},
		{
			"ints",
			1, g.Coalesce(0, 1, 2, 3),
		},
		{
			"ints first",
			1, g.Coalesce(1, 2, 3),
		},
		{
			"ints last",
			1, g.Coalesce(0, 0, 0, 0, 1),
		},
		{
			"ints all zero",
			0, g.Coalesce(0, 0, 0, 0),
		},
		{
			"ints no args",
			0, g.Coalesce[int](),
		},
		{
			"pointers",
			p1, g.Coalesce(nil, p1, p2),
		},
		{
			"pointers first",
			p1, g.Coalesce(p1, p2),
		},
		{
			"pointers last",
			p1, g.Coalesce(nil, nil, p1),
		},
		{
			"pointers all zero",
			(*int)(nil), g.Coalesce[*int](nil, nil, nil),
		},
		{
			"pointers no args",
			(*int)(nil), g.Coalesce[*int](),
		},
	}

	for _, c := range cases {
		if c.exp != c.got {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, c.got)
		}
	}
}

func TestDeref(t *testing.T) {
	cases := []struct {
		name     string
		exp, got any
	}{
		{
			"*int",
			1, g.Deref(g.Ptr(1)),
		},
		{
			"*int nil",
			0, g.Deref[int](nil),
		},
		{
			"*int default",
			2, g.Deref[int](nil, 2),
		},
		{
			"*int not needing default",
			1, g.Deref[int](g.Ptr(1), 2),
		},
		{
			"*string",
			"1", g.Deref(g.Ptr("1")),
		},
		{
			"*string nil",
			"", g.Deref[string](nil),
		},
		{
			"*string default",
			"2", g.Deref[string](nil, "2"),
		},
		{
			"*string not needing default default",
			"1", g.Deref[string](g.Ptr("1"), "2"),
		},
	}

	for _, c := range cases {
		if c.exp != c.got {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, c.got)
		}
	}
}
