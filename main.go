package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func HandleRequest(ctx *fasthttp.RequestCtx) (string, error) {
	r := router.New()
	r.GET("/status", Status)
	return "OK", nil
}

func main() {
	lambda.Start(HandleRequest)
}

func Status(ctx *fasthttp.RequestCtx) {
	addCrossOriginsHeaders(ctx)
	ctx.SetBodyString("Successfull Start")
}

func addCrossOriginsHeaders(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add("Access-Control-Allow-Headers", "*")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE")
}
