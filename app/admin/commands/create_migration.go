package commands

import (
	"fmt"

	"github.com/oussama4/commentify/base/migrate"
)

type CreateMigration struct {
	path string
}

func NewCreateMigration(path string) *CreateMigration {
	return &CreateMigration{path: path}
}

func (cm *CreateMigration) Synopsis() string {
	return "create a database migration file"
}

func (cm *CreateMigration) Help() string {
	return `	create:migration	name`
}

func (cm *CreateMigration) Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a name for migration")
	}
	return migrate.CreateMigration(cm.path, args[0])
}
