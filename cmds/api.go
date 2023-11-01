package cmds

import (
	"fmt"
	"os"
	"test_package/helpers"
	"test_package/modules/user"

	core "github.com/Leakageonthelamp/go-leakage-core"
)

func APIRun() {
	env := core.NewENVPath(helpers.RootDir())

	postgres, err := core.NewDatabase(env.Config()).Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error_postgres: %v\n", err)
		os.Exit(1)
	}

	e := core.NewHTTPServer(&core.HTTPContextOptions{
		ContextOptions: &core.ContextOptions{
			DB:  postgres,
			ENV: env,
		},
	})

	apiGroup := e.Group("/api")

	user.NewUserHTTP(apiGroup)

	core.StartHTTPServer(e, env)
}
