package handler

import (
	"github.com/bruno5200/CSM/group/application"
)

type groupHandler struct {
	GroupService application.Grouper
}

func NewGroupHandler(service application.Grouper) *groupHandler {
	return &groupHandler{GroupService: service}
}
