set GOOS=windows
set GOARCH=amd64
go build -o server.exe main.go server.go user.go
