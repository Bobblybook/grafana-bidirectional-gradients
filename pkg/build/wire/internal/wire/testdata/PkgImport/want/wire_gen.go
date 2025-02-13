// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"example.com/bar"
)

import (
	_ "example.com/anon1"
	_ "example.com/anon2"
)

// Injectors from wire.go:

func injectFooBar() FooBar {
	foo := provideFoo()
	barBar := bar.ProvideBar()
	fooBar := provideFooBar(foo, barBar)
	return fooBar
}
