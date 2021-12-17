package echo

const apiV1 = "/api/v1"

func (r *httpImpl) setRouting() {
	r.public.GET(apiV1+"/ping", r.handler.Ping)
	r.user.GET(apiV1+"/token/refresh/:id", r.handler.refreshToken) //user id

	//account routes
	r.public.POST(apiV1+"/login", r.handler.login)
	r.public.POST(apiV1+"/register", r.handler.register)

	//ticket routes
	r.user.POST(apiV1+"/ticket", r.handler.createNewTicket)
	r.user.PUT(apiV1+"/ticket", r.handler.replyToTicketMessage)
	r.user.GET(apiV1+"/tickets", r.handler.ticketList)
}
