Recover
-

[![CircleCI](https://circleci.com/gh/thepkg/recover.svg?style=svg)](https://circleci.com/gh/thepkg/recover)
[![Go Report Card](https://goreportcard.com/badge/github.com/thepkg/recover)](https://goreportcard.com/report/github.com/thepkg/recover)
[![godoc](https://godoc.org/github.com/thepkg/recover?status.svg)](https://godoc.org/github.com/thepkg/recover)

Package `recover` contains basic functions
which helps to work with `recover` in bit pleasant way.

## Installation

`go get -u github.com/thepkg/recover`

## Usage

````go
import "github.com/thepkg/recover"

func f1() {
  defer recover.All(func(err interface{}) {
    fmt.Printf("got error: %s", err.(error))
  })
}

func f2() {
  defer recover.Any([]error{ErrorUsernameBlank, ErrorUsernameAlreadyTaken}, func(err interface{}) {
    fmt.Printf("got error: %s", err.(error))
  })
}
````
