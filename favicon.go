package middleware

import (
	"github.com/nathanfaucett/ctx"
	"github.com/nathanfaucett/debugger"
	"path"
	"runtime"
)

type FaviconOptions struct {
	Pathname string
}

func Favicon(options *FaviconOptions) func(*ctx.Request, *ctx.Response, func(error)) {
	_, filename, _, _ := runtime.Caller(1)
	dirname := path.Dir(filename)

	debug := debugger.Debug("Favicon")

	if options == nil {
		options = &FaviconOptions{}
	}
	if options.Pathname == "" {
		options.Pathname = "./app/assets/img/favicon.ico"
	}

	pathname := path.Join(dirname, options.Pathname)
	debug.Log("using Favicon " + options.Pathname)

	return func(req *ctx.Request, res *ctx.Response, next func(error)) {
		method := req.Method
		url := req.URL.Path

		if url != "/favicon.ico" || (method != "GET" && method != "HEAD") {
			next(nil)
			return
		}

		if method == "HEAD" {
			debug.Log("Serving HEAD request for " + options.Pathname)
		} else {
			debug.Log("Serving " + options.Pathname)
		}
		res.SendFile(pathname)
		next(nil)
	}
}
