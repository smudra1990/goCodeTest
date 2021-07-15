package vodka

import (
	"encoding/json"
	"net/http"
)

// Context is the most important part of vodka. It allows us to pass variables between middleware.
type Context struct {
	Request    *http.Request
	Writer     http.ResponseWriter
	PathParams map[string]string
}

//JSON response.
func (c *Context) JSON(code int, obj interface{}) {
	c.Writer.WriteHeader(code)

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error()))
	}

	c.Writer.Header().Add("content-type", "application/json")
	c.Writer.Write(jsonBytes)
}
