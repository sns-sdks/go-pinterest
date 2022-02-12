package main

import (
	"fmt"
	"github.com/sns-sdks/go-pinterest/pinterest"
	"os"
)

var (
	AppID     = os.Getenv("app_id")
	AppSecret = os.Getenv("app_secret")
)

func main() {
	app := pinterest.NewAuthorizationAPP(pinterest.AuthorizationAPP{
		ClientID:     AppID,
		ClientSecret: AppSecret,
		RedirectURI:  "https://localhost/",
		Scope:        "pins:read,user_accounts:read",
	})
	fmt.Println(app.String())

	authUrl := app.GetAuthorizationURL()
	fmt.Println(authUrl)

	fmt.Println("Enter Code: ")
	var code string
	//"https://localhost/?code=b26f3de00c9d923a5b3a116716628b557e1c9854&state=go-pinterest"
	fmt.Scanln(&code)
	token, err := app.GenerateAccessToken(code)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)

	pin := app.GetUserClient()
	u, _ := pin.UserAccount.GetUserAccount("")
	fmt.Println(u)
}
