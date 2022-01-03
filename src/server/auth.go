package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"regexp"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	abcs "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"
	"github.com/xvello/letsblockit/src/db"
)

const (
	authPagePrefix = "/auth"
	pageContextKey = "pageContext"
)

func (s *Server) setupAuth() error {
	session, err := buildSessionState(s.store, s.options.RootUrl)
	if err != nil {
		return err
	}

	s.auth = authboss.New()
	s.auth.Config.Paths.Mount = authPagePrefix
	s.auth.Config.Paths.RootURL = s.options.RootUrl.String()
	s.auth.Config.Storage.Server = db.NewUserStore(s.store)
	s.auth.Config.Storage.SessionState = session
	s.auth.Config.Core.ViewRenderer = &authRenderer{pages: s.pages}
	s.auth.Config.Modules.RegisterPreserveFields = []string{defaults.FormValueEmail}

	defaults.SetCore(&s.auth.Config, false, false)
	s.auth.Config.Core.BodyReader = buildBodyReader()
	return s.auth.Init()
}

func (s *Server) buildAuthHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		hc := s.buildPageContext(c, "")
		r := c.Request().WithContext(context.WithValue(c.Request().Context(), pageContextKey, hc))
		http.StripPrefix(authPagePrefix, s.auth.Config.Core.Router).ServeHTTP(c.Response(), r)
		return nil
	}
}

// buildSessionState returns a cookie store with keys persisted in the db.
// If no keys are found (first start), new ones are generated and saved.
func buildSessionState(store db.Store, rootUrl url.URL) (*abcs.SessionStorer, error) {
	var keys [][]byte
	if err := store.RunCtxTx(context.Background(), func(ctx context.Context, q db.Querier) error {
		ek, e := q.GetCookieKeys(ctx)
		switch e {
		case nil:
			keys = append(keys, ek.HashKey, ek.BlockKey)
			return nil
		case db.NotFound:
			nk := db.SetCookieKeysParams{
				HashKey:  securecookie.GenerateRandomKey(32),
				BlockKey: securecookie.GenerateRandomKey(32),
			}
			keys = append(keys, nk.HashKey, nk.BlockKey)
			return q.SetCookieKeys(ctx, nk)
		default:
			return e
		}
	}); err != nil {
		return nil, err
	}

	cookieStore := sessions.NewCookieStore(keys...)
	cookieStore.MaxAge(int((time.Hour * 24 * 7) / time.Second)) // One week
	cookieStore.Options.HttpOnly = true
	cookieStore.Options.Domain = rootUrl.Host
	cookieStore.Options.Secure = rootUrl.Scheme == "https"
	cookieStore.Options.Path = rootUrl.Path

	return &abcs.SessionStorer{
		Name:  "session",
		Store: cookieStore,
	}, nil
}

func buildBodyReader() *defaults.HTTPBodyReader {
	emailRule := defaults.Rules{
		FieldName:  defaults.FormValueEmail,
		Required:   true,
		MatchError: "Must be a valid e-mail address",
		MustMatch:  regexp.MustCompile(`.*@.*\.[a-z]+`),
	}
	passwordRule := defaults.Rules{
		FieldName: defaults.FormValuePassword,
		MinLength: 8,
	}

	return &defaults.HTTPBodyReader{
		UseUsername: false,
		ReadJSON:    false,
		Rulesets: map[string][]defaults.Rules{
			"login":         {emailRule},
			"register":      {emailRule, passwordRule},
			"confirm":       {defaults.Rules{FieldName: defaults.FormValueConfirm, Required: true}},
			"recover_start": {emailRule},
			"recover_end":   {passwordRule},
		},
		Confirms: map[string][]string{
			"register":    {defaults.FormValuePassword, authboss.ConfirmPrefix + defaults.FormValuePassword},
			"recover_end": {defaults.FormValuePassword, authboss.ConfirmPrefix + defaults.FormValuePassword},
		},
		Whitelist: map[string][]string{
			"register": {defaults.FormValueEmail, defaults.FormValuePassword},
		},
	}
}
