package middleware

import (
	"github.com/nathanfaucett/debugger"
	"github.com/nathanfaucett/rest"
	"os"
	"path"
)

type FaviconOptions struct {
	Pathname string
}

func Favicon(options *FaviconOptions) func(*rest.Request, *rest.Response, func(error)) {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	debug := debugger.Debug("Favicon")

	if options == nil {
		options = &FaviconOptions{}
	}
	if options.Pathname == "" {
		options.Pathname = "./app/assets/img/favicon.ico"
	}

	pathname := path.Join(dirname, options.Pathname)
	debug.Log("using Favicon " + options.Pathname)

	return func(req *rest.Request, res *rest.Response, next func(error)) {
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
