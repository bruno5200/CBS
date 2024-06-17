package presenter

import d "github.com/bruno5200/CSM/group/domain"

func GroupSuccessResponse(group *d.Group) map[string]interface{} {
	return map[string]interface{}{
		"group":   group,
		"success": true,
	}
}

func GroupsSuccessResponse(groups *[]d.Group) map[string]interface{} {
	return map[string]interface{}{
		"groups":  groups,
		"success": true,
	}
}

func GroupErrorResponse(err error) map[string]interface{} {
	return map[string]interface{}{
		"error":   err.Error(),
		"success": false,
	}
}

func GroupDisableResponse() map[string]interface{} {
	return map[string]interface{}{
		"success": true,
	}
}
