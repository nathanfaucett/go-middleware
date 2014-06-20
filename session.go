package middleware

import (
	"github.com/nathanfaucett/debugger"
	"github.com/nathanfaucett/rest"
	"github.com/nathanfaucett/util"
	"net/http"
	"strconv"
	//"crypto/sha256"
	//"encoding/hex"
	//"fmt"
)

type Session struct {
	id     string
	cookie *http.Cookie
	values *map[interface{}]interface{}
}

func NewSession() *Session {
	this := new(Session)
	this.values = new(map[interface{}]interface{})
	return this
}
func (this *Session) Id() string {
	return this.id
}
func (this *Session) Cookie() *http.Cookie {
	return this.cookie
}
func (this *Session) Values() *map[interface{}]interface{} {
	return this.values
}

type SessionsCookieOptions struct {
	Path     string
	MaxAge   int
	Domain   string
	HttpOnly bool
	Secure   bool
}
type SessionsOptions struct {
	Key             string
	Path            string
	TrustProxy      bool
	RollingSessions bool
	Secret          string
	CookieOptions   *SessionsCookieOptions
}

func Sessions(options *SessionsOptions) func(*rest.Request, *rest.Response, func(error)) {
	debug := debugger.Debug("Sessions")
	if options == nil {
		options = &SessionsOptions{
			CookieOptions: &SessionsCookieOptions{},
		}
	}
	if options.Key == "" {
		options.Key = "Rest.sid"
	}
	if options.Path == "" {
		options.Path = "/"
	}
	if options.Secret == "" {
		options.Secret = util.Prng(24)
	}

	key := options.Key
	path := options.Path
	trustProxy := options.TrustProxy
	rollingSessions := options.RollingSessions
	secret := options.Secret

	debug.Log(
		"using Sessions with Options" +
			"\n\tkey: " + key +
			"\n\tpath: " + path +
			"\n\ttrustProxy: " + strconv.FormatBool(trustProxy) +
			"\n\trollingSessions: " + strconv.FormatBool(rollingSessions) +
			"\n\tsecret: " + secret)

	return func(req *rest.Request, res *rest.Response, next func(error)) {
		//cookie, err := req.Cookie(key)

		next(nil)
	}
}
