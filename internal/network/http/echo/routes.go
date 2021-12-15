package echo

func (r *httpImpl) setRouting() {
	r.echo.POST("api/v1/login", r.handler.login)
	r.echo.POST("api/v1/register", r.handler.register)
}
