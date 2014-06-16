package middleware

import (
	"github.com/nathanfaucett/debugger"
	"github.com/nathanfaucett/rest"
	"runtime"
	"path"
)

type FaviconOptions struct {
	Pathname string
}

func Favicon(options *FaviconOptions) func(req *rest.Request, res *rest.Response, next func(err error)) {
	_, filename, _, _ := runtime.Caller(1)
	dirname := path.Dir(filename)
	
	debug := debugger.Debug("Favicon")
	
	if (options == nil) {
		options = &FaviconOptions{}
	}
	if options.Pathname == "" {
		options.Pathname = "./app/assets/img/favicon.ico"
	}
	
	pathname := path.Join(dirname, options.Pathname)
	
	return func(req *rest.Request, res *rest.Response, next func(err error)) {
		method := req.Method
        url := req.URL.Path
		
		if (url != "/favicon.ico" || (method != "GET" && method != "HEAD")) {
			next(nil)
            return
		}
		
		debug("Serving "+ options.Pathname)
		res.SendFile(pathname)
		next(nil)
	}
}