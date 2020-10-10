package engine

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/novliang/general/response"
	"net/http"
	"time"
)

func DefaultHttpErrorHandler(err error, c echo.Context) {

	var code = http.StatusInternalServerError
	var msg string

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if h, o := he.Message.(string); o {
			msg = h
		} else {
			msg = "Internal Server Error !"
		}
	} else {
		msg = http.StatusText(code)
	}

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			err := c.NoContent(code)
			if err != nil {
				c.Logger().Error(err)
			}
		} else {
			r := new(response.Response)
			r.Message = msg
			r.Code = code
			err := c.JSON(200, r)
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}
}

type EchoConfig struct {
	Middleware       []echo.MiddlewareFunc
	HttpErrorHandler echo.HTTPErrorHandler
	Validator        echo.Validator
	Logger           echo.Logger
}

type Echo struct {
	*echo.Echo
}

type Context struct {
	echo.Context
	RequestTime int64
}

func (c *Context) Out(i interface{}) error {
	//New Encoder
	e := json.NewEncoder(c.Context.Response())

	//Get Header
	header := c.Response().Header()
	if header.Get(echo.HeaderContentType) == "" {
		header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	}

	//Set Header Code
	c.Response().WriteHeader(http.StatusOK)

	//Appointment Response
	ar := new(response.Response)
	ar.Code = 0
	ar.Message = "success"
	ar.Data = i

	return e.Encode(ar)
}

func (e *Echo) EngineName() string {
	return "echo"
}

func NewEcho(config EchoConfig) *Echo {
	e := Echo{
		echo.New(),
	}

	for _, middleware := range config.Middleware {
		e.Use(middleware)
	}

	//Extending
	e.Use(func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			think := &Context{Context: c, RequestTime: time.Now().Unix()}
			return handlerFunc(think)
		}
	})

	if config.Validator != nil {
		e.Validator = config.Validator
	}

	if config.HttpErrorHandler != nil {
		e.HTTPErrorHandler = config.HttpErrorHandler
	} else {
		e.HTTPErrorHandler = DefaultHttpErrorHandler
	}

	return &e
}
