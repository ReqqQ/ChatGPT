package DomainModelChatGPT

import (
	"encoding/json"

	"github.com/ReqqQ/ChatGPT/Infrastructure/Base"
	InfrastructureChatGPTAPI "github.com/ReqqQ/ChatGPT/Infrastructure/ChatGPTAPI"
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
