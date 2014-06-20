package middleware

import (
	"github.com/nathanfaucett/debugger"
	"github.com/nathanfaucett/rest"
	"os"
	"path"
	"strings"
)

type StaticServeOptions struct {
	Root      string
	Directory string
	Index     string
}

func StaticServe(options *StaticServeOptions) func(*rest.Request, *rest.Response, func(error)) {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	debug := debugger.Debug("StaticServe")

	if options == nil {
		options = &StaticServeOptions{
			Index: "index.html",
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
	serveIndex := options.Index != ""
	index := path.Join(directory, options.Index)
	debug.Log("using StaticServe from " + options.Directory + " root " + root)

	return func(req *rest.Request, res *rest.Response, next func(error)) {
		method := req.Method

		if method != "GET" && method != "HEAD" {
			next(nil)
			return
		}

		url := req.URL.Path

		if url == "/" && serveIndex {
			debug.Log("Serving " + options.Index)
			res.SendFile(index)
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
			debug.Log("Serving GET request for " + url)
		}
		url = string(url[len(root):len(url)])
		fileName := path.Join(directory, url)

		res.SendFile(fileName)
		next(nil)
	}
}
