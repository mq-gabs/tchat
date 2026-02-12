# :computer: tchat

Chat app via terminal. It's composed by a client and server. You need to run the server first and then run the client and configurate the client to access it.

## :mag_right: How it works

### Server

- Just build and run
- There's no configuration needed
- There's no persistence implement (yet), it saves all stuff in-memory

### Client

- It's a CLI REPL (own CLI loop)
- Needs to set server host
- There's no configuration file implemented (yet), it saves all in-memory

## :runner: How to run

### Server

1. Build

```bash
go build -o tchat-server tchat.com/server
```

2. Run

```bash
./tchat-server
```

### Client

1. Build

```bash
go build -o tchat tchat.com/client
```

2. Start the CLI REPL

```bash
./tchat
```

3. Set host

```bash
welcome!
your tchat client is not configured...
please, type the server host

host:
```

4. Set username

```bash
now type your username
user name:
```

5. You'll be registered in the server

```bash
logged successfully!
id: <id>
name: <name>

tchat >
```

6. Check your data running `whoami`
7. Start chat with someone else registered

```bash
chat -userid=<id>
```

## :rocket: Roadmap

- Implement persistence with SQLite
- Implement client configuration file
- Add friends list to configuration file
- Add different servers config to configuration file
