package middleware

import (
	"github.com/nathanfaucett/debugger"
	"github.com/nathanfaucett/rest"
	"github.com/nathanfaucett/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MemeroyStore struct {
}

type Session struct {
	id     string
	cookie *http.Cookie
	values map[interface{}]interface{}
	isNew  bool
	store  rest.Store
}

func NewSession(isNew bool) *Session {
	this := new(Session)

	this.values = make(map[interface{}]interface{})
	this.isNew = isNew

	if isNew {
		this.id = util.Prng(24)
	}

	return this
}

func (this *Session) Id() string {
	return this.id
}
func (this *Session) Cookie() *http.Cookie {
	return this.cookie
}
func (this *Session) Values() map[interface{}]interface{} {
	return this.values
}
func (this *Session) IsNew() bool {
	return this.isNew
}
func (this *Session) Store() rest.Store {
	return this.store
}

type SessionsOptions struct {
	Name       string
	Path       string
	MaxAge     int
	Domain     string
	HttpOnly   bool
	Secure     bool
	TrustProxy bool
	Secret     string
}

func Sessions(options *SessionsOptions) func(*rest.Request, *rest.Response, func(error)) {
	debug := debugger.Debug("Sessions")

	if options == nil {
		options = &SessionsOptions{}
	}
	if options.Name == "" {
		options.Name = "Rest.sid"
	}
	if options.Path == "" {
		options.Path = "/"
	}
	if options.MaxAge == 0 {
		options.MaxAge = 3600
	}
	if options.Secret == "" {
		options.Secret = util.Prng(24)
	}

	debug.Log(
		"using Sessions with Options" +
			"\n\tname: " + options.Name +
			"\n\tpath: " + options.Path +
			"\n\ttrustProxy: " + strconv.FormatBool(options.TrustProxy) +
			"\n\tsecret: " + options.Secret)

	return func(req *rest.Request, res *rest.Response, next func(error)) {
		if strings.Index(req.URL.Path, options.Path) != 0 {
			next(nil)
			return
		}

		var session *Session
		cookie, err := req.Cookie(options.Name)
		if err == nil {
			session = NewSession(false)
			session.id = util.Unsign(cookie.Value, options.Secret)
			session.cookie = cookie
		} else {
			cookie = &http.Cookie{
				Name:     options.Name,
				Path:     options.Path,
				Domain:   options.Domain,
				MaxAge:   options.MaxAge,
				Expires:  time.Now().Add(time.Duration(time.Duration(options.MaxAge) * time.Second)),
				Secure:   options.Secure,
				HttpOnly: options.HttpOnly,
			}
			session = NewSession(true)
			session.cookie = cookie
		}
		req.Session = session

		res.On("header", func() {
			if !session.isNew {
				debug.Log("Already set Cookie")
				return
			}

			cookie.Value = util.Sign(session.id, options.Secret)
			serialized := cookie.String()
			debug.Log("Setting new Cookie " + serialized)
			res.SetHeader("Set-Cookie", serialized)
		})

		next(nil)
	}
}
