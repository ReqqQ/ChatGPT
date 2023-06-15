package AppChatGPT

import (
	"encoding/json"

	DomainModelChatGPT "github.com/ReqqQ/ChatGPT/DomainModel/ChatGPT"
)

type AppChatGPTFactory struct {
}

type ChatGPTDTO struct {
	Json json.RawMessage
}
type MessageDTO struct {
	Content string `json:"content"`
}

func (c ChatGPTDTO) GetJson() json.RawMessage {
	return c.Json
}
func (factory AppChatGPTFactory) GetMessageDTO(message DomainModelChatGPT.Message) MessageDTO {
	return MessageDTO{Content: message.GetContent()}
}
