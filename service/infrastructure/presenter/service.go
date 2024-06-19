package presenter

import d "github.com/bruno5200/CSM/service/domain"

func ServiceSuccessResponse(service *d.Service) map[string]interface{} {
	return map[string]interface{}{
		"service": service,
		"success": true,
	}
}

func ServicesSuccessResponse(services *[]d.Service) map[string]interface{} {
	return map[string]interface{}{
		"services": services,
		"success":  true,
	}
}

func ServiceErrorResponse(err error) map[string]interface{} {
	return map[string]interface{}{
		"error":   err.Error(),
		"success": false,
	}
}

func ServiceDisableResponse() map[string]interface{} {
	return map[string]interface{}{
		"success": true,
	}
}
