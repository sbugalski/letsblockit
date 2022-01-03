package server

import (
	"context"
	"net/url"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	abcs "github.com/volatiletech/authboss-clientstate"
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"
	"github.com/xvello/letsblockit/src/db"
)

const authPrefix = "/auth"

func buildAuth(store db.Store, rootUrl url.URL) (*authboss.Authboss, error) {
	session, err := buildSessionState(store, rootUrl)
	if err != nil {
		return nil, err
	}

	ab := authboss.New()
	ab.Config.Paths.Mount = authPrefix
	ab.Config.Paths.RootURL = rootUrl.String()
	ab.Config.Storage.Server = db.NewUserStore(store)
	ab.Config.Storage.SessionState = session
	ab.Config.Core.ViewRenderer = defaults.JSONRenderer{}

	defaults.SetCore(&ab.Config, false, false)
	if err := ab.Init(); err != nil {
		return nil, err
	}
	return ab, nil
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
