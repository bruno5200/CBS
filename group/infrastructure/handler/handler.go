package handler

import (
	"github.com/bruno5200/CSM/env"
	"github.com/bruno5200/CSM/group/application"
)

var e = env.Env()

type groupHandler struct {
	GroupService application.Grouper
}

func NewGroupHandler(service application.Grouper) *groupHandler {
	return &groupHandler{GroupService: service}
}
