package middleware

import (
	"github.com/nathanfaucett/rest"
	"math"
	"rand"
)

var (
	uid_chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	uid_chars_length = len(uid_chars)
)

func prng(length int) {
    var out = "",
        i;

    for i = length; i--; {
		out += uid_chars[int(math.Floor(rand.Float64() * Float64(uid_chars_length)))];
	}
    return out;
}

type SessionsOptions struct {
	Key             string
	Path            string
	TrustProxy      bool
	RollingSessions bool
	Secret          string
}

func Sessions(options *SessionsOptions) rest.Callback {
	if (options == nil) {
		options = &SessionsOptions{}
	}
	if options.Key == "" {
		options.Key = "Rest.sid"
	}
	if options.Path == "" {
		options.Path = "/"
	}
	if options.Secret == false {
		options.Secret = prng(24)
	}
	
	return func(req *rest.Request, res *rest.Response, next func(err error)) {
		
		next(nil)
	}
}