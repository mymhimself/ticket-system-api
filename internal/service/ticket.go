package service

type Ticket interface {
	SendInterfaceByUser()
	LoadInterfaceByAdmin()
	LoadInterfaceByUser()
}
