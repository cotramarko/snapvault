# Snapvault 
![snapvault](snapvault.webp)
A PostgreSQL snapshot tool that effortlessly captures and restores precise snapshots of your database.
___
**⚠️ Note:** This tool is designed for use during development and should not be used in production.
## 📸 Why Snapvault?
The snapvault CLI tool is intended to be used during local development as an easy way to capture and restore snapshots of the database, making it possible to quickly restore the database to a previous state. 

It uses the [template](https://www.postgresql.org/docs/current/manage-ag-templatedbs.html) functionality in Postgres to create clones of databases, which is faster than using `pg_dump`/`pg_restore`. This means that all clones are actually stored as separate databases on the same Postgres server as the original database.

It supports basic commands such as `save`, `restore`, `list` and `delete`:

```shell
$ snapvault save <snapshot_name> 
$ snapvault restore <snapshot_name>
$ snapvault list
$ snapvault delete <snapshot_name>
```

Snapvault is similar to projects like [DLSR](https://github.com/mixxorz/DSLR) and [Stellar](https://github.com/fastmonkeys/stellar). However, unlike those projects snapvault is written in Go and delivered as a standalone binary, making it possible to use the tool without having to rely on Python or managing any other dependencies.  

## ⚙️ Installation
Binaries are available in both Intel and ARM versions for OSX/Darwin, Linux and Windows and can be found under the [Releases](https://github.com/cotramarko/snapvault/releases) section.

### Using Homebrew
```shell
$ brew tap cotramarko/tools
$ brew install snapvault
```

### Manual Download

**OSX/Darwin & Linux**
```shell
# Change binary depending on your platform
# Linux: 
#   snapvault_Linux_arm64
#   snapvault_Linux_i386
#   snapvault_Linux_x86_64
# OSX:
#   snapvault_Darwin_arm64
#   snapvault_Darwin_x86_64

# Bash:
$ TARGET=snapvault_Darwin_x86_64
# Fish:
$ set TARGET snapvault_Darwin_x86_64
```
```shell
# Download binary and make it executable
$ sudo curl -fsSL -o /usr/local/bin/snapvault https://github.com/cotramarko/snapvault/releases/latest/download/$TARGET
$ sudo chmod +x /usr/local/bin/snapvault
```

**Windows**

[Download and unzip the archive](https://github.com/cotramarko/snapvault/releases) for your target platform. The unzipped archive contains an `.exe` of the binary.  

## 🔧 How to Use Snapvault

### Providing Database URL
The database URL can be provided in multiple ways

#### Option 1 - Using a `snapvault.toml` file
Create a `snapvault.toml` file containing the database URL 
```toml
url = "postgres://acmeuser:acmepassword@localhost:5432/acmedb"
```
The URL will be used whenever `snapvault` is in the same directory as the file.

#### Option 2 - Setting the `DATABASE_URL` environment variable
Another option is to set the `DATABASE_URL` environment variable to the database URL
```shell
# Bash:
$ export DATABASE_URL=postgres://acmeuser:acmepassword@localhost:5432/acmedb
# Fish:
$ set -x DATABASE_URL postgres://acmeuser:acmepassword@localhost:5432/acmedb
```
If both the `DATABASE_URL` is set and a `snapvault.toml` file is present then the `snapvault.toml` file will be preferred.   

#### Option 3 - Passing it explicitly with `--url` flag
Another option is to explicitly pass the database URL with the `--url` flag whenever running a snapvault command
```shell
$ snapvault list --url=postgres://acmeuser:acmepassword@localhost:5432/acmedb
```
The `--url` flag will always override any of the other ways of specifying the URL.

### Basic Commands
```shell
$ snapvault save fix/foobar
Created snapshot fix/foobar

$ snapvault restore fix/foobar
Restored snapshot fix/foobar

$ snapvault list
╭────────────┬──────────────────────┬─────────┬─────────╮
│ NAME       │        CREATED       │   SIZE  │ COMMENT │
├────────────┼──────────────────────┼─────────┼─────────┤
│ fix/foobar │ 2025-08-02T14:12:18Z │ 7561 kB │         │
╰────────────┴──────────────────────┴─────────┴─────────╯

$ snapvault delete fix/foobar
Deleted snapshot fix/foobar
```
