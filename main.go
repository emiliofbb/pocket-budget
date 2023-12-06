package main

import (
    "log"
    "os"
    "strings"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"
    "github.com/pocketbase/pocketbase/plugins/migratecmd"

    _ "pocket-budget/migrations"
)

func main() {
    app := pocketbase.New()

    // check if it was executed using go run
    isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

    migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
        Automigrate: isGoRun,
    })

    // serves static files from the provided public dir (if exists)
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.GET("/static/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), true),
    ) 
        return nil
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}
