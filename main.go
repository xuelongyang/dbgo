package main

import (
	"dbgo/constent"
	"dbgo/database/mysql"
	"dbgo/database/oracle"
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
		mysql.MySQL()
	case Oracle:
		oracle.Oracle()
	default:
		fmt.Println("Database not supported, currently supported input \"mysql\" and \"oracle\"!")
	}
}

func main() {
	args := os.Args
	if len(args) > 2 && args[1] == "-t" {
		matchingDatabase(args[2])
		os.Args = append(os.Args[:1], os.Args[2:]...)
	} else if len(args) == 2 && args[1] == "-v" {
		fmt.Println("Dbgo version " + constent.Version + " for " + runtime.GOOS + "/" + runtime.GOARCH)
		fmt.Println("Author: xuelongyang")
		fmt.Println("Github: https://github.com/xuelongyang/dbgo")
	} else {
		fmt.Print(constent.Usage)
	}
}
