package middleware

import (
	"github.com/nathanfaucett/ctx"
	"github.com/nathanfaucett/debugger"
	"path"
	"runtime"
	"strings"
)

type StaticServeOptions struct {
	Root       string
	Directory  string
	ServeIndex bool
}

func StaticServe(options *StaticServeOptions) func(*ctx.Request, *ctx.Response, func(error)) {
	_, filename, _, _ := runtime.Caller(1)
	dirname := path.Dir(filename)

	debug := debugger.Debug("StaticServe")

	if options == nil {
		options = &StaticServeOptions{
			ServeIndex: true,
		}
	}
	if options.Root == "" {
		options.Root = "/assets/"
	}
	if string(options.Root[0]) != "/" {
		options.Root = "/" + options.Root
	}
	if string(options.Root[len(options.Root)-1]) != "/" {
		options.Root += "/"
	}
	if options.Directory == "" {
		options.Directory = "./app/assets/"
	}

	root := options.Root
	directory := path.Join(dirname, options.Directory)
	serveIndex := options.ServeIndex
	debug.Log("using StaticServe from " + options.Directory + " root " + root)

	return func(req *ctx.Request, res *ctx.Response, next func(error)) {
		method := req.Method
		url := req.URL.Path

		if url == "/" && serveIndex {
			fileName := path.Join(directory, "index.html")
			res.SendFile(fileName)
			next(nil)
			return
		}

		if method != "GET" && method != "HEAD" {
			next(nil)
			return
		}
		if strings.Index(url, root) != 0 {
			next(nil)
			return
		}

		if method == "HEAD" {
			debug.Log("Serving HEAD request for " + url)
		} else {
			debug.Log("Serving " + url)
		}
		url = string(url[len(root):len(url)])
		fileName := path.Join(directory, url)

		res.SendFile(fileName)
		next(nil)
	}
}
