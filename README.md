# Dbgo - The Database Command Line Interacting Tool
![Version](https://img.shields.io/badge/version-1.0.0-orange)
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

This is a command line tool for interacting with the Dbgo database

### Download and Install

You can download the binary executable we publish and choose the executable that suits your system environment to 
download. You can run executable files directly in your environment for database connection operations.

Currently supported databases are:

| Database | Version   |
| --- |-----------|
|MySQL| 5.7 and later |

We currently support the following runtime environments:

| OS | Arch |
|---|---|
| Windows | amd64 |
| Linux | amd64 |
| Linux | arm64 |

### Usage

Dbgo is a tool for database various relational databases.

1. **MySQL**

    you can use the following command to connect to the database:
    ```shell
    dbgo -t mysql -h your_host -P your_port -u your_username -p
    ```
    The following are the parameters supported by the command:

        -h  Connect to host, default value is "localhost"
        -P  Port number to use for connection, default value is "3306"
        -u  User for login, default value is "root"
        -p  Password to use when connecting to server.It is ciphertext input.
