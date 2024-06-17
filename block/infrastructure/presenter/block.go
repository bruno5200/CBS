package presenter

import d "github.com/bruno5200/CSM/block/domain"

func BlockSuccessResponse(block *d.Block) map[string]interface{} {
	return map[string]interface{}{
		"block":   block,
		"success": true,
	}
}

func BlocksSuccessResponse(blocks *[]d.Block) map[string]interface{} {
	return map[string]interface{}{
		"blocks":  blocks,
		"success": true,
	}
}

func BlockCreateResponse(url string) map[string]interface{} {
	return map[string]interface{}{
		"url":     url,
		"success": true,
	}
}

func BlockErrorResponse(err error) map[string]interface{} {
	return map[string]interface{}{
		"error":   err.Error(),
		"success": false,
	}
}

func BlockDisableResponse() map[string]interface{} {
	return map[string]interface{}{
		"success": true,
	}
}
