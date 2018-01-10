package main

import (
	"fmt"
	"strconv"

	"github.com/valyala/fasthttp"
)

// Item show determine location information
func LocationShow(ctx *fasthttp.RequestCtx) {

	id, err := strconv.ParseInt(ctx.UserValue("id").(string), 10, 64)
	if err == nil {
		fmt.Fprintf(ctx, "locations[%v] = %+v ", id, locations[id])
	} else {
		fmt.Fprintf(ctx, "Value can't parse to int64")
	}

}
