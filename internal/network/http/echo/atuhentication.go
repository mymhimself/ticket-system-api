package echo

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) refreshToken(c echo.Context) error {
	//idString := c.Param("id")
	//id, err := strconv.Atoi(idString)

	//if err != nil {
	//	h.logger.Error(err.Error())
	//	return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest})
	//}

	//token, err := h.accountService.Login(c.Request().Context(), req.Username, req.Password)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, response.Error{Error: err.Error(), Status: http.StatusBadRequest}) // i know can better error handling
	//}
	//
	//return c.JSON(http.StatusOK, response.Login{Token: token, Status: http.StatusOK})
	return nil
}
