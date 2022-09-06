package app

import (
	"Q2Bank/app/user"
	"Q2Bank/store"
)

// ContainerApp is a struct to export instances of apps
type ContainerApp struct {
	User user.App
}

// Options is a struct to receive options to apps
type Options struct {
	Stores *store.ContainerStore
}

// NewContainerApp create a new instance of apps
func NewContainerApp(opts Options) *ContainerApp {
	return &ContainerApp{
		User: user.NewApp(opts.Stores),
	}
}
