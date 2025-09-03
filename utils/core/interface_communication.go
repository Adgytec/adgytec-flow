package core

type Communicaiton interface {
	SendMail([]string, string) error
}
