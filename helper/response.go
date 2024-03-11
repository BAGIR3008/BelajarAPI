package helper

func ResponseFormat(status int, message any, data ...map[string]any) (int, map[string]any) {
	result := map[string]any{
		"code":    status,
		"message": message,
	}

	for _, part := range data {
		for key, value := range part {
			result[key] = value
		}
	}
	return status, result
}
