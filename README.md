# rpi-status-cli
CLI program like a shell. You can check the status of the board based on various parameters.

### Compile

```bash
GOOS=linux GOARCH=arm64 go build -o rpi-shell main.go setup.go
```

### Launch

```bash
chmod +x rpi-shell
./rpi-shell
```