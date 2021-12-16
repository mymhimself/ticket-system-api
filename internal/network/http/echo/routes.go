package echo

func (r *httpImpl) setRouting() {
	r.public.GET("/api/v1/ping", r.handler.Ping)
	r.public.POST("/api/v1/login", r.handler.login)
	r.public.POST("/api/v1/register", r.handler.register)

	r.user.GET("/api/v1/token/refresh/:id", r.handler.refreshToken) //user id
	r.user.POST("/api/v1/ticket", r.handler.newTicket)
}
