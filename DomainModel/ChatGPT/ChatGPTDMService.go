package DomainModelChatGPT

import (
	"encoding/json"

	"ChatGPT/Infrastructure/Base"
	InfrastructureChatGPTAPI "ChatGPT/Infrastructure/ChatGPTAPI"
)

type ChatGPTService struct {
	BaseRequest Base.BaseRequest
	ChatGPTApi  InfrastructureChatGPTAPI.ChatGPTAPIRepository
}

func (cg ChatGPTService) TranslateIntoHumanLanguage(json json.RawMessage) Message {
	cg.BaseRequest.Init()
	response := cg.ChatGPTApi.TranslateIntoHumanLanguage(json, cg.BaseRequest)

	return getChatGPTMessageResponse(response)
}
