package cmds

import (
	"fmt"
	"os"
	"test_package/helpers"
	"test_package/models"

	core "github.com/Leakageonthelamp/go-leakage-core"
)

func MigrateRun() {
	env := core.NewENVPath(helpers.RootDir())

	db, err := core.NewDatabase(env.Config()).Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error_postgres: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(
		&models.User{},
	)
}
