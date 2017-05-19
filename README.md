## token-svc

All this does is listen on a custom TCP port and respond to HTTP GETs with random 40-character length token strings.

```sh
go build token-svc.go
./token-svc 1337
```

Now, with Docker:

```sh
make build
make image
make run
```

Or, with a custom port on Docker:

```sh
docker run -it --rm -p 1337:1337 token-svc 1337
```
