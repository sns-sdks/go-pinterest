# go-pinterest

A simple Go wrapper for  Pinterest REST API (Beta) (5.x) ✨ 🍰 ✨

[![Test Status](https://github.com/sns-sdks/go-pinterest/workflows/tests/badge.svg)](https://github.com/sns-sdks/go-pinterest/actions?query=workflow%3Atests)
[![Test Coverage](https://codecov.io/gh/sns-sdks/go-pinterest/branch/main/graph/badge.svg)](https://codecov.io/gh/sns-sdks/go-pinterest)
[![Go Report Card](https://goreportcard.com/badge/github.com/sns-sdks/go-pinterest)](https://goreportcard.com/report/github.com/sns-sdks/go-pinterest)

## Installation

```shell
# Go Modules
require github.com/sns-sdks/go-pinterest
```

## Usage

### Authentication

You can initial the client with access token.

```go
client := pinterest.NewBearerClient("Your bearer token")
u, err := client.UserAccount.GetUserAccount("")
fmt.Println(u, err)
```

Or you can give oauth flow by hand, You can follow the [`authorize example`](https://github.com/sns-sdks/go-pinterest/blob/main/example/authentication/main.go) 

More usage detail see the [`Example`](https://github.com/sns-sdks/go-pinterest/blob/main/example)

## Features

- OAuth
- UserAccount
- Boards
- Pins
- Media
- AdAccounts
