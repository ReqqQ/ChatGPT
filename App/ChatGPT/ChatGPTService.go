package AppChatGPT

import DomainModelChatGPT "github.com/ReqqQ/ChatGPT/DomainModel/ChatGPT"

type AppChatGPTService struct {
	DomainChatGPTService DomainModelChatGPT.ChatGPTService
	AppChatGPTFactory    AppChatGPTFactory
}

func (gpt AppChatGPTService) GetDataFromChatGPT(dto ChatGPTDTO) MessageDTO {
	messageVO := gpt.DomainChatGPTService.TranslateIntoHumanLanguage(dto.GetJson())

	return gpt.AppChatGPTFactory.GetMessageDTO(messageVO)
}
