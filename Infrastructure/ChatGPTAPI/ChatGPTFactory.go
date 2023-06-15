package InfrastructureChatGPTAPI

import "encoding/json"

const BASE_MODEL = "gpt-3.5-turbo"
const REQUEST_TYPE_POST = "POST"

type RequestBody struct {
	Model       string     `json:"model"`
	Messages    []Messages `json:"messages"`
	Temperature float32    `json:"temperature"`
	N           int        `json:"n"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func getRequestBody(json json.RawMessage) RequestBody {
	return RequestBody{
		Model: BASE_MODEL,
		Messages: []Messages{
			{Role: "user", Content: string(json)},
		},
		Temperature: 1,
		N:           1,
	}
}
