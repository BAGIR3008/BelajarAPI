package tools

func Response(status int, message string, data ...map[string]any) (int, map[string]any) {
	json := map[string]any{
		"code":    status,
		"message": message,
	}

	for _, part := range data {
		for key, value := range part {
			json[key] = value
		}
	}
	return status, json
}
