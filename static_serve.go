package middleware

import (
	"github.com/nathanfaucett/debugger"
	"github.com/nathanfaucett/rest"
	"runtime"
	"strings"
	"path"
)

type StaticServeOptions struct {
	Root       string
	Directory  string
	ServeIndex bool
}

func StaticServe(options *StaticServeOptions) rest.Callback {
	_, filename, _, _ := runtime.Caller(1)
	dirname := path.Dir(filename)
	
	debug := debugger.Debug("StaticServe")
	
	if (options == nil) {
		options = &StaticServeOptions{
			ServeIndex: true,
		}
	}
	if options.Root == "" {
		options.Root = "/assets/"
	}
	if options.Directory == "" {
		options.Directory = "./app/assets/"
	}
	
	root := options.Root
	directory := path.Join(dirname, options.Directory)
	serveIndex := options.ServeIndex;
	
	return func(req *rest.Request, res *rest.Response, next func(err error)) {
		method := req.Method
        url := req.URL.Path
		
		if url == "/" && serveIndex {
			fileName := path.Join(directory, "index.html")
			res.SendFile(fileName)
			next(nil)
			return
		}
		
        if (method != "GET" && method != "HEAD") {
            next(nil)
            return
        }
        if (strings.Index(url, root) != 0) {
            next(nil)
            return
        }

		debug("Serving "+ url)
		url = string(url[len(root):len(url)])
		fileName := path.Join(directory, url)
		
		res.SendFile(fileName)
		next(nil)
	}
}