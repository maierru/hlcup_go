package locations

import "github.com/valyala/fasthttp"
import "fmt"
import "strconv"

// Item show determine location information
func Item(ctx *fasthttp.RequestCtx) {
	value, e := ctx.UserValue("id").(string)
	if e {
		fmt.Fprintf(ctx, "Value %v type assertion with error", value)
	}
	id, err := strconv.ParseInt(value, 10, 64)
	if err == nil {
		fmt.Fprintf(ctx, "Value %v can't parse to int64 as %T", id)
	} else {
		fmt.Fprintf(ctx, "%v is int64", id)
	}

}
