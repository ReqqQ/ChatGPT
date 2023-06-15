package InfrastructureChatGPTAPI

import (
	"encoding/json"

	"ChatGPT/Infrastructure/Base"
)

const CHATGPT_COMPLETIONS_ENDPOINT = "https://api.openai.com/v1/chat/completions"

type ChatGPTAPIRepository struct{}

func (api ChatGPTAPIRepository) TranslateIntoHumanLanguage(jsonMessage json.RawMessage, baseRequest Base.BaseRequest) []byte {
	request, response := baseRequest.BaseSettings()
	buildCompletionsRequestData(request)
	request.SetBody(transformRequestBodyToJson(jsonMessage))
	baseRequest.CallAPI(request, response)

	return response.Body()
}
