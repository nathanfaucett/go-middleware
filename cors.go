package middleware

import (
	"github.com/nathanfaucett/debugger"
	"github.com/nathanfaucett/rest"
	"strconv"
)

type CorsOptions struct {
	Origin      string
	Methods     string
	Credentials string
	MaxAge      int
	Headers     string
}

func Cors(options *CorsOptions) func(*rest.Request, *rest.Response, func(error)) {
	debug := debugger.Debug("Cors")

	if options == nil {
		options = &CorsOptions{}
	}
	if options.Origin == "" {
		options.Origin = "*"
	}
	if options.Methods == "" {
		options.Methods = "GET,POST,PUT,PATCH,HEAD,OPTIONS,DELETE"
	}
	if options.Credentials == "" {
		options.Credentials = "false"
	}
	if options.Headers == "" {
		options.Headers = ""
	}

	origin := options.Origin
	methods := options.Methods
	credentials := options.Credentials
	var maxAge string
	if options.MaxAge == 0 {
		maxAge = ""
	} else {
		maxAge = strconv.Itoa(options.MaxAge)
	}
	corsHeaders := options.Headers

	debug.Log(
		"using Cors with Options" +
			"\n\torigin: " + origin +
			"\n\tmethods: " + methods +
			"\n\tcredentials: " + credentials +
			"\n\tmaxAge: " + maxAge +
			"\n\tcorsHeaders: " + corsHeaders)

	return func(req *rest.Request, res *rest.Response, next func(error)) {
		var headers string
		if corsHeaders == "" {
			headers = req.GetHeader("Access-Control-Request-Headers")
		} else {
			headers = corsHeaders
		}

		if origin != "" {
			res.SetHeader("Access-Control-Allow-Origin", origin)
		}
		if methods != "" {
			res.SetHeader("Access-Control-Allow-Methods", methods)
		}
		if credentials != "" {
			res.SetHeader("Access-Control-Allow-Credentials", credentials)
		}
		if headers != "" {
			res.SetHeader("Access-Control-Allow-Headers", headers)
		}
		if maxAge != "" {
			res.SetHeader("Access-Control-Allow-Max-Age", maxAge)
		}

		next(nil)
	}
}
