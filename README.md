# Dbgo - The Database Command Line Interacting Tool
![Version](https://img.shields.io/badge/version-1.1.0-orange)
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

This is a command line tool for interacting with the Dbgo database

### Download and Install

You can download the binary executable we publish and choose the executable that suits your system environment to 
download. You can run executable files directly in your environment for database connection operations.

Currently supported databases are:

| Database   | Version        |
|------------|----------------|
| MySQL      | 5.7 and later  |
| Oracle     | 10.2 and later |
| PostgreSQL | 10 and later   |

We currently support the following runtime environments:

| OS | Arch |
|---|---|
| Windows | amd64 |
| Linux | amd64 |
| Linux | arm64 |

### Usage

Dbgo is a tool for database various relational databases.

1. **MySQL**

    You can use the following command to connect to the database:
    ```shell
    dbgo -t mysql -h your_host -P your_port -u your_username -p
    ```
    The following are the parameters supported by the command:

        -h  Connect to host, default value is "localhost"
        -P  Port number to use for connection, default value is "3306"
        -u  User for login, default value is "root"
        -p  Password to use when connecting to server.It is ciphertext input.
2. **Oracle**

    You can use the following command to connect to the database:
    ```shell
    dbgo -t oracle -h your_host -P your_port -u your_username -S your_service -p
    ```
    The following are the parameters supported by the command:

        -h  Connect to host, default value is "localhost"
        -P  Port number to use for connection, default value is "1521"
        -S  Service name to use for connection, default value is "orcl"
        -u  User for login, default value is "SYS"
        -p  Password to use when connecting to server. It is ciphertext input.

3. **PostgreSQL**

    You can use the following command to connect to the database:
    ```shell
    dbgo -t postgresql -h your_host -P your_port -u your_username -d database_name -p
    ```
    The following are the parameters supported by the command:

        -h  Connect to host, default value is "localhost"
        -P  Port number to use for connection, default value is "5432"
        -u  User for login, default value is "postgres"
        -p  Password to use when connecting to server. It is ciphertext input.
        -d  The database to be used, default value is "postgres".
        -s  Whether or not to use SSL, default value is "disable".