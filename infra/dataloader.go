package infra

import (
	"encoding/json"
	"io"
)

type Importer[T any] interface {
	Import() (*T, error)
}

type JsonImporter[T any] struct {
	reader io.Reader
}

func (i *JsonImporter[T]) Import() (*T, error) {
	res := new(T)
	if err := json.NewDecoder(i.reader).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}
