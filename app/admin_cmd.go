package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/oussama4/commentify/store"
)

type adminCmd struct {
	synopsis string
	fs       *flag.FlagSet
	store    store.Store
}

func newAdminCmd(store store.Store) *adminCmd {
	return &adminCmd{
		synopsis: "create an admin user",
		fs:       flag.NewFlagSet("admin", flag.PanicOnError),
		store:    store,
	}
}

func (ac *adminCmd) Synopsis() string {
	return ac.synopsis
}

func (ac *adminCmd) Help() string {
	u := `usage: commentify admin
		create an admin user
		`

	return u
}

func (ac *adminCmd) Init(args []string) {
	ac.fs.Parse(args)
}

func (ac *adminCmd) Run() error {
	name, email := "", ""
	fmt.Print("name: ")
	fmt.Fscan(os.Stdin, &name)
	fmt.Print("email: ")
	fmt.Fscan(os.Stdin, &email)

	userId, err := ac.store.CreateUser(name, email)
	if err != nil {
		return err
	}

	fmt.Printf("admin user created succesfully with id: %s", userId)
	return nil
}
