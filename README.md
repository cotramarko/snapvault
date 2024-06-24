# Snapvault 
![snapvault](snapvault.webp)
A PostgreSQL backup tool that effortlessly captures and restores precise snapshots of your database.
___
**⚠️ Note:** This tool is designed for use during development and should not be used in production.
## 📸 Why Snapvault?
The snapvault CLI tool is intended to be used during local development as an easy way to capture and restore snapshots of the database, making it possible to quickly restore the database to a previous state. It uses the [template](https://www.postgresql.org/docs/current/manage-ag-templatedbs.html) functionality in Postgres to create clones of databases, which is faster than using `pg_dump`/`pg_restore`. It supports basic commands such as `save`, `restore`, `list` and `delete`:

```shell
$ snapvault save <snapshot_name> 
$ snapvault restore <snapshot_name>
$ snapvault list
$ snapvault delete <snapshot_name>
```

Snapvault is similar to projects like [DLSR](https://github.com/mixxorz/DSLR) and [Stellar](https://github.com/fastmonkeys/stellar). However, unlike those projects snapvault is written in Go and delivered as a standalone binary, making it possible to use the tool without having to rely on Python or managing any other dependencies. 

## ⚙️ Installation
Binaries are available in both Intel and ARM versions for OSX/Darwin, Linux and Windows and can be found under the [Releases](https://github.com/cotramarko/snapvault/releases) section.

### Manual Download
```shell
# Change binary depending on your platform
$ TARGET=snapvault_Darwin_x86_64
$ sudo curl -fsSL -o /usr/local/bin/snapvault https://github.com/cotramarko/snapvault/releases/latest/download/$TARGET
$ sudo chmod +x /usr/local/bin/snapvault
```
### Using `brew`
```shell
$ brew tap cotramarko/tools
$ brew install snapvault
```

## 🔧 How to Use Snapvault

### Specifying Database
The database URL can be specified in multiple ways. Either by a `snapvault.toml` file
(containing `url=<connection-string>`), or by setting the environment variable
`$DATABASE_URL=<connection-string>`, or by passing it as a flag via `--url=<connection-string>`.

The `--url` flag will always override any of the other ways of specifying the URL. If both a
`snapvault.toml` file is present and `$DATABASE_URL` is set, then the `snapvault.toml` file will be prioritised.

### Basic Commands
```shell
$ snapvault save fix/foobar
Created snapshot fix/foobar

$ snapvault restore fix/foobar
Restored snapshot fix/foobar

$ snapvault list
╭────────────┬──────────────────────┬─────────╮
│ NAME       │        CREATED       │    SIZE │
├────────────┼──────────────────────┼─────────┤
│ fix/foobar │ 2024-06-23T15:37:39Z │ 7561 kB │
╰────────────┴──────────────────────┴─────────╯

$ snapvault delete fix/foobar
Deleted snapshot fix/foobar
```
