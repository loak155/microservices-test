```bash:
$ docker-compose up

$ go run main.go

$ grpcurl -plaintext localhost:8080 list
$ grpcurl -plaintext localhost:8080 list user.UserService
$ grpcurl -plaintext -d '{"username": "test_user", "email": "test@example.com", "password": "password"}' localhost:8080 user.UserService.Signup
$ grpcurl -plaintext -d @ localhost:8080 user.UserService.Signup
  {"username": "test_user", "email": "test@example.com", "password": "password"}
  Ctrl+Z (UnixはCtrl+D)
$ grpcurl -plaintext -d @ localhost:8080 user.UserService.Login
  {"email": "test@example.com", "password": "password"}
  Ctrl+Z (UnixはCtrl+D)

```
