package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStorage(t *testing.T) {
	a := assert.New(t)
	err := Store.Set("/tmp", 12)
	a.Nil(err)
	var value int
	err = Store.Get("/tmp", &value)
	a.Nil(err)
	a.Equal(12, value)
}

func TestNewStorage2(t *testing.T) {
	a := assert.New(t)
	err := Store.Set("/mc/a", 12)
	a.Nil(err)
	var value int
	err = Store.Get("/mc/a", &value)
	a.Nil(err)
	a.Equal(12, value)
}

func TestStorage_Exists(t *testing.T) {
	a := assert.New(t)
	a.Equal(true,Store.Exists("/tmp"))
	a.Equal(false,Store.Exists("/mc"))
	a.Equal(false,Store.Exists("/00"))

}

func TestPathExists(t *testing.T) {
	a := assert.New(t)
	a.Equal(PathExists("/mm"), false)
	a.Equal(PathExists("_"), false)
}
