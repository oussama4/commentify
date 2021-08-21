package app

import (
	"flag"
	"fmt"
	"os"

	"github.com/oussama4/commentify/store"
)

type adminCmd struct {
	name  string
	fs    *flag.FlagSet
	store store.Store
}

func newAdminCmd(store store.Store) *adminCmd {
	return &adminCmd{
		name:  "admin",
		fs:    flag.NewFlagSet("admin", flag.PanicOnError),
		store: store,
	}
}

func (ac *adminCmd) Name() string {
	return ac.name
}

func (ac *adminCmd) Usage() string {
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

	userId, err := ac.store.CreateUser(name, email, 1)
	if err != nil {
		return err
	}

	fmt.Printf("admin user created succesfully with id: %s", userId)
	return nil
}
