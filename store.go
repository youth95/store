package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Codec interface {
	Encoder(v interface{}) ([]byte, error)
	Decoder(p []byte, v interface{}) error
}

type Storage struct {
	basePath string
	codec    Codec
}

type JSONCodec int

func (*JSONCodec) Encoder(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (*JSONCodec) Decoder(p []byte, v interface{}) error {
	return json.Unmarshal(p, v)
}

var Store *Storage

func init() {
	s, _ := os.UserHomeDir()
	Store = NewStorage(filepath.Join(s, ".go-store"), new(JSONCodec))
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func NewStorage(basePath string, codec Codec) *Storage {
	if !PathExists(basePath) {
		err := os.Mkdir(basePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return &Storage{basePath, codec}
}

func (storage *Storage) Get(path string, v interface{}) error {
	savePath := filepath.Join(storage.basePath, path)
	p, err := ioutil.ReadFile(savePath)
	if err != nil {
		return err
	}
	return storage.codec.Decoder(p, v)
}

func (storage *Storage) Set(path string, v interface{}) error {
	savePath := filepath.Join(storage.basePath, path)
	dir, _ := filepath.Split(savePath)
	if !PathExists(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	p, err := storage.codec.Encoder(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(savePath, p, 0666)
}

func (storage *Storage) Exists(path string) bool {
	savePath := filepath.Join(storage.basePath, path)
	stat, err := os.Stat(savePath)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}
