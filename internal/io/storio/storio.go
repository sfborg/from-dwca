package storio

import "github.com/sfborg/from-dwca/internal/ent/stor"

type storio struct {
}

func New() stor.Storage {
	return &storio{}
}
