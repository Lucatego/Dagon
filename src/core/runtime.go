package core

import (
	"os"
	"os/user"
)

/*
	TODO: Get the dir of work, get the username, get the hostname.
*/

var (
	Hostname string
	Username string
	Dir      string
)

func StartRuntime() {
	// Get the hostname
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	Hostname = host
	// Get the Username
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	Username = user.Username
	// Get the current working directory
	Dir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}
