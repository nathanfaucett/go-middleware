package middleware

import (
	"github.com/nathanfaucett/debugger"
	"github.com/nathanfaucett/rest"
	"net/http"
	"math"
	"math/rand"
	//"crypto/sha256"
	//"encoding/hex"
	//"strconv"
	//"fmt"
)

var (
	uid_chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

func prng(length int) string {
    out := ""
    for i := 0; i < length; i++ {
		out += string(uid_chars[int(math.Floor(rand.Float64() * 62))]);
	}
    return out;
}

type Session struct {
	id               string
	cookie          *http.Cookie
	values          *map[interface{}]interface{}
}

func NewSession() *Session{
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

func Sessions(options *SessionsOptions) func(req *rest.Request, res *rest.Response, next func(err error)) {
	debug := debugger.Debug("Sessions")
	if (options == nil) {
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
		options.Secret = prng(24)
	}
	
	//key := options.Key
	//path := options.Path
	//trustProxy := options.TrustProxy
	//rollingSessions := options.RollingSessions
	//secret := options.Secret
	
	debug.Log("using Sessions")
	
	return func(req *rest.Request, res *rest.Response, next func(err error)) {
		//cookie, err := req.Cookie(key)
		
		req.Session = NewSession()
		next(nil)
	}
}