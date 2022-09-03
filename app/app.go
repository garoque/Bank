package app

import "Q2Bank/store"

// ContainerApp is a struct to export instances of apps
type ContainerApp struct {
}

// Options is a struct to receive options to apps
type Options struct {
	Stores *store.ContainerStore
}

// NewContainerApp create a new instance of apps
func NewContainerApp(opts Options) *ContainerApp {
	return &ContainerApp{}
}
