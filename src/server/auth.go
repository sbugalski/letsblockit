package server

import (
	"github.com/volatiletech/authboss/v3"
	"github.com/volatiletech/authboss/v3/defaults"
	"github.com/xvello/letsblockit/src/db"
)

const authPrefix = "/authboss"

func buildAuth(store db.Store) (*authboss.Authboss, error) {
	ab := authboss.New()
	ab.Config.Storage.Server = db.NewUserStore(store)
	ab.Config.Paths.Mount = authPrefix
	ab.Config.Paths.RootURL = "https://letsblock.it/"

	defaults.SetCore(&ab.Config, false, false)
	if err := ab.Init(); err != nil {
		return nil, err
	}
	return ab, nil
}
