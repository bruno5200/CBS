package presenter

import (
	"time"

	d "github.com/bruno5200/CSM/group/domain"
)

type GroupResponse struct {
	Id          string  `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	ServiceId   string  `json:"serviceId,omitempty"`
	ServiceName string  `json:"serviceName,omitempty"`
	Active      bool    `json:"active"`
	Blocks      []Block `json:"blocks,omitempty"`
}

type Block struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Checksum    string    `json:"checksum,omitempty"`
	Extension   string    `json:"extension"`
	Url         string    `json:"url,omitempty"`
	UploadedAt  time.Time `json:"uploadedAt,omitempty"`
	GroupId     string    `json:"groupId"`
	GroupName   string    `json:"groupName,omitempty"`
	ServiceId   string    `json:"serviceId"`
	ServiceName string    `json:"serviceName,omitempty"`
	Active      bool      `json:"active"`
}

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
