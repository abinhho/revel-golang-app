package controllers

import "github.com/revel/revel"

type ApiController struct {
	*revel.Controller
	callBack string
}

type Response struct {
	Code    int         `json:"code"`
	Results interface{} `json:"results,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/* api response code */
const (
	OK int = iota
	WARN_NOT_FOUND
	ERR_VALIDATE
	ERR_FATAL
)

func (c *ApiController) ErrorMessage(s int) string {

	switch s {
		case ERR_VALIDATE:
			return "Error Validate"
		default:
			return ""
	}
}
func (c *ApiController) Response(s interface{}) revel.Result {

	if c.callBack != "" {
		return c.RenderJSONP(c.callBack, s)
	} else {
		return c.RenderJSON(s)
	}
}
