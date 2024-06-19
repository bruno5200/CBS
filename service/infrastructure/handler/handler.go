package handler

import "github.com/bruno5200/CSM/service/application"

type serviceHandler struct {
	serviceService application.Serviceer
}

func NewServiceHandler(serviceService application.Serviceer) *serviceHandler {
	return &serviceHandler{serviceService: serviceService}
}
