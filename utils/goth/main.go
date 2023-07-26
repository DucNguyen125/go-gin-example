package goth

import (
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

func Init() {
	goth.UseProviders(
		facebook.New(
			os.Getenv("FACEBOOK_CLIENT_ID"),
			os.Getenv("FACEBOOK_SECRET"),
			os.Getenv("BASE_URL")+"/v1/auth/facebook/callback",
		),
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_SECRET"),
			os.Getenv("BASE_URL")+"/v1/auth/google/callback",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		),
	)
	store := sessions.NewCookieStore([]byte(os.Getenv("SECRET_SESSION_KEY")))
	gothic.Store = store
}
