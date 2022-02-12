package main

import (
	"fmt"
	"github.com/sns-sdks/go-pinterest/pinterest"
	"os"
)

var bearerToken = os.Getenv("access_token")

func main() {
	pin := pinterest.NewBearerClient(bearerToken)

	u, _ := pin.UserAccount.GetUserAccount("")
	fmt.Println(u)
}
