package main

import (
	"github.com/lakevilladom/goreq"
	"github.com/lakevilladom/gospider"
)

func main() {
	s := gospider.NewSpider()
	s.SeedTask(
		goreq.Get("https://httpbin.org/get"),
		func(ctx *gospider.Context) {
			ctx.AddItem(ctx.Resp.Text)
		},
	)
	s.OnItem(func(ctx *gospider.Context, i interface{}) interface{} {
		ctx.Println(i)
		return i
	})
	s.Wait()
}
