package delivery

import (
	"net/http"
	"nusatech/features/users"
	"nusatech/utils/middlewares"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var validate = validator.New()

type userHandler struct {
	srv users.Service
}

func New(e *echo.Echo, srv users.Service) {
	handler := userHandler{srv: srv}
	e.POST("/user", handler.Create())
	e.POST("/login", handler.Login())
	e.GET("/user", handler.ShowAll())
	e.PUT("/user", handler.Update(), middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
}

func (uh *userHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		er := validate.Struct(input)
		if er != nil {
			if strings.Contains(er.Error(), "min") {
				return c.JSON(http.StatusBadRequest, FailResponse("min. 4 character"))
			} else if strings.Contains(er.Error(), "max") {
				return c.JSON(http.StatusBadRequest, FailResponse("max. 30 character"))
			} else if strings.Contains(er.Error(), "email") {
				return c.JSON(http.StatusBadRequest, FailResponse("must input valid email"))
			}
			return c.JSON(http.StatusBadRequest, FailResponse(er.Error()))
		}

		cnv := ToCore(input)
		res, err := uh.srv.Create(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return c.JSON(http.StatusBadRequest, FailResponse("duplicate email on database"))
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("cannot encrypt password"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server."))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success register user", ToResponse(res, "user")))
	}
}

func (uh *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToCore(input)
		res, err := uh.srv.Login(cnv)
		if err != nil {
			if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, FailResponse("password not match."))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusAccepted, SuccessResponse("success login", ToResponse(res, "login")))
	}
}

func (uh *userHandler) ShowAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uh.srv.ShowAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusOK, SuccessResponse("success get all data", ToResponse(res, "getall")))
	}
}

func (uh *userHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		IdUser := middlewares.ExtractToken(c)
		if IdUser == 0 {
			return c.JSON(http.StatusUnauthorized, FailResponse("cannot validate token"))
		}

		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("cannot bind input"))
		}

		cnv := ToCore(input)
		res, err := uh.srv.Update(cnv, IdUser, input.OldEmail)
		if err != nil {
			if strings.Contains(err.Error(), "email") {
				return c.JSON(http.StatusBadRequest, FailResponse("email not match"))
			} else if strings.Contains(err.Error(), "found") {
				return c.JSON(http.StatusBadRequest, FailResponse("must input old email"))
			}
			return c.JSON(http.StatusInternalServerError, FailResponse("there is problem on server"))
		}

		return c.JSON(http.StatusAccepted, SuccessResponse("success update user", ToResponse(res, "user")))
	}
}
