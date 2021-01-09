package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "log"
)

type H map[string]interface{}

type Context struct {

	Writer http.ResponseWriter
	Req *http.Request

	Path string
	Method string

	StatusCode int

}

func newContext(rw http.ResponseWriter, req *http.Request) *Context {

	return &Context{
		Writer: rw,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}

}

func (context *Context) PostForm(key string) string {

	return context.Req.FormValue(key)

}

func (context *Context) Query(key string) string {

	return context.Req.URL.Query().Get(key)

}

func (context *Context) SetStatus(code int) {

	context.StatusCode = code
	context.Writer.WriteHeader(code)

}

func (context *Context) SetHeader(key, val string){

	context.Writer.Header().Set(key, val)

}

const ContentType = "Content-Type"

func (context *Context) WriteString(code int, format string, value ...interface{}) {

	context.SetHeader(ContentType, "text/plain")
	context.SetStatus(code)
	context.Writer.Write([]byte(fmt.Sprintf(format, value...)))

}

func (context *Context) WriteJson(code int, obj interface{}) {

	context.SetHeader(ContentType, "application/json")
	context.SetStatus(code)
	encoder := json.NewEncoder(context.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Writer, err.Error(), 500)
	}

}

func (context *Context) Data(code int, data []byte) {

	context.SetStatus(code)
	context.Writer.Write(data)

}

func (context *Context) HTML(code int, html string) {

	context.SetHeader(ContentType, "text/html")
	context.SetStatus(code)
	context.Writer.Write([]byte(html))

}
