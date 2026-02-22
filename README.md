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

1. Build Client and Server
> The binaries will be saved on `dist/`
```bash
make all
```


2. Run Server

```bash
./tchat-server
```

3. Start the CLI REPL

```bash
./tchat
```

4. Set username

```bash
now type your username
user name:
```

5. Register the server you just started
> The default port is `8080`

```bash
server add -host:<host:port>
```

6. Check your data running `whoami`

7. Start chat with someone else registered on the server you connected

```bash
chat -userid=<id>
```

