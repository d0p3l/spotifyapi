package main

import (
	"net/http"
	"time"

	"github.com/d0p3l/spotifyapi/internal/app/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)




func main() {
	e := echo.New()

	// e.Use(authMiddleWare)
	e.GET("/login", completeAuth)
	e.GET("/", userInfo)
	e.GET("/authurl", getAuthUrl)

	e.Logger.Fatal(e.Start(":8080"))
}

func userInfo(ctx echo.Context) error {
	tok := giveToken(ctx)

	// // cookie, err := ctx.Cookie("access_token")
	// // if err != nil {
	// // 	return err
	// // }

	client := spotify.New(auth.Client(ctx.Request().Context(), tok))
	user, err := client.CurrentUser(ctx.Request().Context())
	if err != nil {
		return err
	}
	// fmt.Fprint(ctx.Response().Writer, user.User)
	return ctx.JSON(200, user.User)
}

func completeAuth(e echo.Context) error {
	tok, err := auth.Token(e.Request().Context(), state, e.Request())
	if err != nil {
		// http.Error(e.Response().Writer, "Couldn't get token", http.StatusForbidden)
		return err
	}
	if st := e.FormValue("state"); st != state {
		// http.NotFound(e.Response().Writer, e.Request())
		// log.Fatalf("State mismatch: %s != %s\n", st, state)
		return err
	}

	return e.JSON(200, tok)
}

// func authMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		tok := giveToken(ctx)

// 		err := next(ctx)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}
// }

// func refreshToken(ctx echo.Context) error {

// }

func giveToken(ctx echo.Context) *oauth2.Token {
	// date, _ := time.Parse(time.RFC1123, ctx.QueryParam("expiry"))
	// ctx.QueryParam("token_type")
	// ctx.QueryParam("refresh_token")

	tok := &oauth2.Token{
		AccessToken:  ctx.Request().Header.Get("Authorization"),
		TokenType:    ctx.FormValue("token_type"),
		RefreshToken: ctx.FormValue("refresh_token"),
		Expiry:       time.Now(), // TODO сделать нормальный time date
	}

	return tok
}

func getAuthUrl(ctx echo.Context) error {
	url := auth.AuthURL(state)
	return ctx.String(http.StatusOK, url)
}
