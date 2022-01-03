package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/volatiletech/authboss/v3"
	"github.com/xvello/letsblockit/src/pages"
)

type (
	authRenderer struct {
		pages PageRenderer
	}

	formField struct {
		Name  string
		Label string
		Type  string
	}
)

var (
	authTabs = []formTab{{
		Title: "Create account",
		Type:  "register",
	}, {
		Title: "Login",
		Type:  "login",
	}, {
		Title: "Recover password",
		Type:  "recover",
	}}

	authPages = map[string]struct {
		Title  string
		Target string
		Tabs   []formTab
		Fields []formField
		Intro  string
	}{
		"login": {
			Title:  "Log into your account",
			Target: "login",
			Tabs:   authTabs,
			Fields: []formField{{
				Name:  "email",
				Label: "E-mail",
				Type:  "email",
			}, {
				Name:  "password",
				Label: "Password",
				Type:  "password",
			}, {
				Name:  "login",
				Label: "Log In",
				Type:  "submit",
			}},
		},
		"register": {
			Title:  "Create a new account",
			Target: "register",
			Tabs:   authTabs,
			Fields: []formField{{
				Name:  "email",
				Label: "E-mail",
				Type:  "email",
			}, {
				Name:  "password",
				Label: "Password",
				Type:  "password",
			}, {
				Name:  "confirm_password",
				Label: "Confirm password",
				Type:  "password",
			}, {
				Name:  "register",
				Label: "Register",
				Type:  "submit",
			}},
		},
		"recover_start": {
			Title:  "Recover your account",
			Target: "recover",
			Tabs:   authTabs,
			Intro:  "Enter your e-mail below to receive a recovery link by e-mail.",
		},
		"recover_end": {
			Title:  "Recover your account",
			Target: "recover/end",
			Tabs:   authTabs,
			Intro:  "Set a new password for your account.",
		},
	}
)

func (a authRenderer) Load(_ ...string) error {
	return nil
}

func (a authRenderer) Render(ctx context.Context, name string, data authboss.HTMLData) (output []byte, contentType string, err error) {
	page, ok := authPages[name]
	if !ok {
		return nil, "", fmt.Errorf("unsupported form %s", name)
	}
	hc, ok := ctx.Value(pageContextKey).(*pages.Context)
	if !ok {
		return nil, "", errors.New("page context missing")
	}

	hc.Title = page.Title
	hc.NoBoost = true
	hc.Add("type", name)
	hc.Add("form", page)
	for k, v := range data {
		hc.Add(k, v)
	}

	b, err := a.pages.RenderRaw("auth-form", hc)
	fmt.Println(hc.Data)
	return b, "text/html", err
}
