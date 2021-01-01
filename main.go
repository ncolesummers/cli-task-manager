package main

import (
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/ncolesummers72/task-manager-cli/cmd"
	"github.com/ncolesummers72/task-manager-cli/db"
)

func main() {
	home, _ := homedir.Dir()

	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
