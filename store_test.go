package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStorage(t *testing.T) {
	a := assert.New(t)
	err := store.Set("/tmp", 12)
	a.Nil(err)
	var value int
	err = store.Get("/tmp", &value)
	a.Nil(err)
	a.Equal(12, value)
}

func TestNewStorage2(t *testing.T) {
	a := assert.New(t)
	err := store.Set("/mc/a", 12)
	a.Nil(err)
	var value int
	err = store.Get("/mc/a", &value)
	a.Nil(err)
	a.Equal(12, value)
}

func TestStorage_Exists(t *testing.T) {
	a := assert.New(t)
	a.Equal(true,store.Exists("/tmp"))
	a.Equal(false,store.Exists("/mc"))
	a.Equal(false,store.Exists("/00"))

}

func TestPathExists(t *testing.T) {
	a := assert.New(t)
	a.Equal(PathExists("/mm"), false)
	a.Equal(PathExists("_"), false)
}
