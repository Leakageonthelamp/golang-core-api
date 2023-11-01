package main

import (
	"test_package/cmds"
	"test_package/consts"

	core "github.com/Leakageonthelamp/go-leakage-core"
)

func main() {
	switch core.NewEnv().Config().Service {
	case consts.ServiceAPI:
		cmds.APIRun()
	case consts.ServiceMigration:
		cmds.MigrateRun()
	default:
		panic("Unknown service exit application")
	}
}
