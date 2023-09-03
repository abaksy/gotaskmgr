/*
Copyright © 2023 Aronya Baksy abaksy@gmail.com
*/
package main

import (
	"github.com/abaksy/gotaskmgr/cmd"
	"github.com/abaksy/gotaskmgr/db"
)

func main() {
	db.InitDB()
	cmd.Execute()
}
