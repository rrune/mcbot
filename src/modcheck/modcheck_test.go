package modcheck

import (
	"testing"

	"github.com/matryer/is"
)

var mc Modcheck

func TestMain(m *testing.M) {
	mc = Init()

	m.Run()
}

func TestCheckMod(t *testing.T) {
	is := is.New(t)
	b := mc.checkMod("442282")
	is.True(b)
}

func TestCheck(t *testing.T) {
	mc.Check()
}

func TestCache(t *testing.T) {
	mc.Cache()
	mc.GetCache()
}
