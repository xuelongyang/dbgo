package constent

// Version info.
const Version = "1.0.0"

// Usage info.
const Usage = `Dbgo is a tool for database various relational databases.
Currently supported databases are:
  MySQL
Usage:
  dbgo -t [database] <command> [arguments]
The commands are:
  -t  database type, must parameter, currently supported input "mysql",the different database parameters are as follows:
        mysql: <-h> [host] <-P> [port] <-u> [username] <-p> [password]
          -h  Connect to host, default value is "localhost"
          -P  Port number to use for connection, default value is "3306"
          -u  User for login, default value is "root"
          -p  Password to use when connecting to server. It is ciphertext input.
  -v    print dbgo version
`
