package echo

import (
	"context"
	"fmt"
	"github.com/mymhimself/ticket-system-api/internal/config"
	"github.com/mymhimself/ticket-system-api/internal/entity/enum"
	"github.com/mymhimself/ticket-system-api/internal/entity/model"
	"github.com/mymhimself/ticket-system-api/internal/network"
	"github.com/mymhimself/ticket-system-api/internal/pkg/logger"
	"github.com/mymhimself/ticket-system-api/internal/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type httpImpl struct { //REST
	echo    *echo.Echo
	handler *handler
	public  *echo.Group
	admin   *echo.Group
	user    *echo.Group
}

func New(cfg config.Authentication, authorizationCfg config.Authorization, logger logger.Logger, accSrv service.Account, ticketSrv service.Ticket, validation service.Validation) network.Rest {
	echoInstance := echo.New()
	//enforcer, err := casbin.NewEnforcer(authorizationCfg.Casbin.ModelPath, authorizationCfg.Casbin.PolicyPath)
	//if err != nil {
	//	log.Fatal("casbin file loader error.\n", err.Error())
	//}
	//echoInstance.Use(casbinMw.Middleware(enforcer))
	echoInstance.Use(middleware.Gzip())
	echoInstance.Use(middleware.RequestID())
	echoInstance.Use(middleware.Recover())
	echoInstance.Use(middleware.Secure())
	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete},
	}))
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte(cfg.Secret),
		Claims:     &model.TokenClaims{},
		ErrorHandler: func(err error) error {
			fmt.Println(err.Error())
			return &echo.HTTPError{
				Code:    http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			}
		},
	}

	public := echoInstance.Group("")
	admin := echoInstance.Group("/admin", middleware.JWTWithConfig(jwtConfig), accessControl(enum.Admin))
	user := echoInstance.Group("/user", middleware.JWTWithConfig(jwtConfig))

	var httpInstance = &httpImpl{
		echo:   echoInstance,
		public: public,
		admin:  admin,
		user:   user,
		handler: &handler{
			logger:                logger,
			accountService:        accSrv,
			ticketService:         ticketSrv,
			authenticationService: nil,
			validationService:     validation,
		}}

	return httpInstance
}

func (r *httpImpl) Start(address string) error {
	r.echo.Use(middleware.Recover())

	r.setRouting()

	return r.echo.Start(address)
}

func (r *httpImpl) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return r.echo.Shutdown(ctx)
}
