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

// Performs recover in case of panic with error ErrorUsernameBlank
// otherwise panic won't be recovered and will be propagated.
defer recover.One(ErrorUsernameBlank, func(err interface{}) {
  fmt.Printf("got error: %s", err)
})

// Performs recover in case of panic with error ErrorUsernameBlank or ErrorUsernameAlreadyTaken
// otherwise panic won't be recovered and will be propagated.
defer recover.Any([]error{ErrorUsernameBlank, ErrorUsernameAlreadyTaken}, func(err interface{}) {
  fmt.Printf("got error: %s", err)
})

// Performs recover for all panics.
defer recover.All(func(err interface{}) {
  fmt.Printf("got error: %s", err)
})
````
