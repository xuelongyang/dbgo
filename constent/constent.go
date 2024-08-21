package constent

// Version info.
const Version = "1.2.0"

// Usage info.
const Usage = `Dbgo is a tool for database various relational databases.
Currently supported databases are:
  MySQL
  Oracle
  PostgreSQL
Usage:
  dbgo -t [database] <command> [arguments]
The commands are:
  -t  database type, must parameter, currently supported input "mysql", "oracle" and "postgresql",the different database parameters are as follows:
        mysql: <-h> [host] <-P> [port] <-u> [username] <-p> [password]
          -h  Connect to host, default value is "localhost"
          -P  Port number to use for connection, default value is "3306"
          -u  User for login, default value is "root"
          -p  Password to use when connecting to server. It is ciphertext input.
        oracle: <-h> [host] <-P> [port] <-S> [service] <-u> [username] <-p> [password]
          -h  Connect to host, default value is "localhost"
          -P  Port number to use for connection, default value is "1521"
          -S  Service name to use for connection, default value is "orcl"
          -u  User for login, default value is "SYS"
          -p  Password to use when connecting to server. It is ciphertext input.
        postgresql: <-h> [host] <-P> [port] <-u> [username] <-p> [password]
          -h  Connect to host, default value is "localhost"
          -P  Port number to use for connection, default value is "5432"
          -u  User for login, default value is "postgres"
          -p  Password to use when connecting to server. It is ciphertext input.
          -d  The database to be used, default value is "postgres".
          -s  Whether or not to use SSL, default value is "disable".
  -v  print dbgo version
`
