package main

import (
	"dbgo/constent"
	"dbgo/database"
	"fmt"
	"os"
	"runtime"
)

// Database Client will support following databases.
const (
	MySQL      = "mysql"      // supported
	Oracle     = "oracle"     // in the future
	PostgreSQL = "postgresql" // in the future
	SqlServer  = "sqlserver"  // in the future
	Sqlite     = "sqlite"     // in the future
)

// Matching database type and create database connection.
func matchingDatabase(dbType string) {
	switch dbType {
	case MySQL:
		database.MySQL()
	default:
		fmt.Println("Database not supported")
	}
}

func main() {
	args := os.Args
	if len(args) > 2 && args[1] == "-t" {
		matchingDatabase(args[2])
	} else if len(args) == 2 && args[1] == "-v" {
		fmt.Println("Dbgo version " + constent.Version + " for " + runtime.GOOS + "/" + runtime.GOARCH)
		fmt.Println("Author: xuelongyang")
	} else {
		fmt.Print(constent.Usage)
	}
}
