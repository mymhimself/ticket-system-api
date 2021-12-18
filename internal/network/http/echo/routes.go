package echo

const apiV1 = "/api/v1"

func (r *httpImpl) setRouting() {
	r.admin.PUT(apiV1+"/ticket", r.handler.replyToTicketMessageByAdmin)
	r.admin.PATCH(apiV1+"/ticket", r.handler.changeTicketStatusByAdmin)
	r.admin.GET(apiV1+"/ticket/:id", r.handler.ticketDetailsByAdmin)
	r.admin.GET(apiV1+"/tickets", r.handler.ticketListByAdmin)

	r.public.GET(apiV1+"/ping", r.handler.Ping)
	r.public.POST(apiV1+"/login", r.handler.login)
	r.public.POST(apiV1+"/register", r.handler.register)
	r.public.GET(apiV1+"/refresh/:id", r.handler.refreshToken) //user id

	//ticket routes
	r.user.POST(apiV1+"/ticket", r.handler.createNewTicket)
	r.user.PUT(apiV1+"/ticket", r.handler.replyToTicketMessage)
	r.user.GET(apiV1+"/ticket/:id", r.handler.ticketDetails)
	r.user.GET(apiV1+"/tickets", r.handler.ticketList)
}
