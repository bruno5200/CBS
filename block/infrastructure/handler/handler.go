package handler

import (
	a "github.com/bruno5200/CSM/block/application"
	"github.com/bruno5200/CSM/env"
)

var e = env.Env()

type blockHandler struct {
	BlockService a.Blocker
}

func NewBlockHandler(service a.Blocker) *blockHandler {
	return &blockHandler{BlockService: service}
}
