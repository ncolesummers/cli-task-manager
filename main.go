package main

import (
	"log"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/ncolesummers72/task-manager-cli/cmd"
	"github.com/ncolesummers72/task-manager-cli/db"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		log.Println(err)
	}

	dbPath := filepath.Join(home, "tasks.db")
	err = db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	cmd.RootCmd.Execute()
}
