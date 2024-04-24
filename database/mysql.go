package database

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

func MySQL() {
	db, err := connect()
	if err != nil {
		fmt.Println(err)
	} else {
		interaction(db)
	}
}

func interaction(db *sql.DB) {
	fmt.Println("MySQL is connected, you may begin execute your commands, input \"exit\" will quit dbgo. ^_^")
	reader := bufio.NewReader(os.Stdin)
	var query string
	for {
		var input string
		var err error
		fmt.Print("mysql> ")
		input, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if runtime.GOOS == "windows" && len(strings.TrimSuffix(input, "\r\n")) == 0 {
			continue
		} else if len(strings.TrimSuffix(input, "\n")) == 0 {
			continue
		}
		if runtime.GOOS == "windows" && input[len(input)-3] == '\\' {
			query += input[:len(input)-3]
			continue
		} else if runtime.GOOS == "linux" && input[len(input)-2] == '\\' {
			query += input[:len(input)-2]
			continue
		}
		// 去除输入中的换行符
		if runtime.GOOS == "windows" {
			query += strings.TrimSuffix(input, "\r\n")
		} else if runtime.GOOS == "linux" {
			query += strings.TrimSuffix(input, "\n")
		}
		// 检查用户是否输入了退出命令
		if strings.TrimSpace(query) == "exit" {
			fmt.Println("Thank you for your use, Bye bye! ^_^")
			os.Exit(0)
		}

		// 执行用户输入的 SQL 查询
		rows, err := db.Query(query)
		query = ""
		if err != nil {
			fmt.Println("Error executing query:", err)
			continue
		}
		defer rows.Close()

		// 获取列名
		columns, err := rows.Columns()
		if err != nil {
			fmt.Println("Error getting columns:", err)
			return
		}
		// 打印列名
		if len(columns) > 0 {
			fmt.Println(strings.Join(columns, "\t"))
			fmt.Println("===========================")
		} else {
			fmt.Println("No columns")
		}

		// 准备一个切片用于保存扫描出的值
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = new(interface{})
		}

		// 遍历查询结果并打印
		for rows.Next() {
			err := rows.Scan(values...)
			if err != nil {
				fmt.Println("Error scanning row:", err)
				return
			}

			// 打印每一行的值
			for _, value := range values {
				if value != nil {
					if reflect.TypeOf(*value.(*interface{})) == reflect.TypeOf(int64(0)) {
						fmt.Printf("%v\t", *value.(*interface{}))
					} else if *value.(*interface{}) == nil {
						fmt.Printf("NULL\t")
					} else {
						fmt.Printf("%s\t", *value.(*interface{}))
					}
				}
			}
			fmt.Println()
		}
	}
}

func connect() (*sql.DB, error) {
	var dbType string
	var host string
	var port uint64
	var username string
	var password string
	var hasP bool
	flag.StringVar(&dbType, "t", "", "database type, must parameter, currently supported input \"mysql\"")
	flag.StringVar(&username, "u", "root", "User for login, default value is \"root\"")
	flag.Uint64Var(&port, "P", 3306, "Port number to use for connection, default value is \"3306\"")
	flag.StringVar(&host, "h", "localhost", "Connect to host, default value is \"localhost\"")
	handlePassword(&hasP)
	if !hasP {
		flag.StringVar(&password, "p", "", "Password to use when connecting to server")
		fmt.Println("")
		flag.Usage()
		os.Exit(0)
	}
	flag.Parse()
	fmt.Print("Enter password: ")
	oldState, err := terminal.GetState(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
	}
	if _, err := terminal.MakeRaw(int(os.Stdin.Fd())); err != nil {
		fmt.Println(err)
	}
	defer func(fd int, oldState *terminal.State) {
		err := terminal.Restore(fd, oldState)
		if err != nil {
			fmt.Println(err)
		}
	}(int(os.Stdin.Fd()), oldState)
	data, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	password = string(data)
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":"+strconv.FormatUint(port, 10)+")/")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}

func handlePassword(hasP *bool) {
	args := os.Args[1:]
	filteredArgs := make([]string, 0, len(args))
	for _, arg := range args {
		if arg != "-p" {
			filteredArgs = append(filteredArgs, arg)
		} else {
			*hasP = true
		}
	}
	os.Args = append([]string{os.Args[0]}, filteredArgs...)
}
