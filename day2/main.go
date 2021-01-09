package main

import(
	"net/http"

	"gee"
)

func main() {
	server := gee.New()
	server.GET("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	server.GET("/hello", func(context *gee.Context) {
		// expect /hello?name=geektutu
		context.WriteString(http.StatusOK, "hello %s, you're at %s\n", context.Query("name"), context.Path)
	})

	server.POST("/login", func(context *gee.Context) {
		context.WriteJson(http.StatusOK, gee.H{
			"username": context.PostForm("username"),
			"password": context.PostForm("password"),
		})
	})

	server.Run(":9999")
}