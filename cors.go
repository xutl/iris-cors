package cors

import (
	"github.com/kataras/iris/context"
)

func New() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			origin := ctx.GetHeader("Origin")
			if len(origin) > 0 {
				ctx.Header("Access-Control-Allow-Origin", origin)
				ctx.Header("Access-Control-Allow-Methods", "GET, OPTIONS ,POST")
				ctx.Header("Access-Control-Allow-Headers", "Content-Type,Content-Encoding")
			}
			if ctx.Method() == "OPTIONS" {
				ctx.Header("Access-Control-Max-Age", "1728000")
			}
		}()

		ctx.Next()
	}
}
