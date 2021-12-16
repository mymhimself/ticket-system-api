package service

type Account interface {
	Login(username, password string) (string, error)
	Register(username, password string) error
	RegisterConfirm(key, phone, otp string) (string, error)
}
