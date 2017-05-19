## token-svc

All this does is listen on a custom TCP port and respond to HTTP GETs with random 40-character length token strings.

```sh
go build token-svc.go
./token-svc 1337
```
