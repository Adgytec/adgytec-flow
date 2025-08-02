package core

type ICommunicaiton interface {
	SendMail([]string, string) error
}
