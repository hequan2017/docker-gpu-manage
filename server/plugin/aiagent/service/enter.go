package service

var Service = new(service)

type service struct {
	Conversation conversation
	Message      message
	Chat         chat
	Config       config
}
