package handlers

import (
	config "github.com/d0p3l/spotifyapi/envconfig"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

const redirectURI = "http://localhost:8080/login"

var (
	confauth = config.New()
	auth     = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate,
		spotifyauth.ScopeUserTopRead, spotifyauth.ScopeUserReadRecentlyPlayed, spotifyauth.ScopeUserLibraryRead),
		spotifyauth.WithClientID(confauth.SPOTIFY_ID), spotifyauth.WithClientSecret(confauth.SPOTIFY_SECRET))
	state = "abc1234"
)
