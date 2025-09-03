package core

type Communication interface {
	SendMail([]string, string) error
}
