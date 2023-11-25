# krampus
Krampus is a replacement for Python's http.server, but with uploading.

### Flags
* --help
* --file-download-path (default "./")
* --file-upload-path (default "./uploads")
* --ssl true|false (default false)
* --ssl-cert-path (default "./cert.pem")
* --ssl-key-path (default "./key.pem")
* --port \<port\> (default 9001)

### SSL
NOTE: If not specificed, requires exactly named 'cert.pem' and 'key.pem' in current directory.
#### Quick Certificate
`openssl req -newkey rsa:2048 -new -nodes -x509 -daexecutedys 3650 -keyout key.pem -out cert.pem`
```
$ ls
cert.pem  key.pem
$ krampus --ssl true --port 8443
krampus(v1.1) starting at port: 8443 (SSL: true)
```

## Uploading

### Windows
`C:\> curl -F "file=@<file>" http[:]//<ip>:<port>/upload`

### Unix
`$ curl -F "file=@<file>" http[:]//<ip>:<port>/upload`


## Downloading

### Windows
`curl -o <file> http[:]//<ip>:<port>/<file>` 

### Unix
`curl -O http[:]//<ip>:<port>/<file>`


## Compiling
NOTE: Requires network connection for the first compile.

### Windows
```
C:\Users\username\krampus> dir
go.mod  go.sum  main.go  README.md
C:\Users\username\krampus> go build -o krampus.exe -ldflags="-s -w"
go.mod  go.sum  krampus.exe  main.go  README.md
```

### Unix
```
$ pwd
/home/username/krampus
$ ls

go.mod  go.sum  main.go  README.md
$ go build -o krampus -ldflags="-s -w"
go.mod  go.sum  krampus  main.go  README.md
```

Note: The script 'krampus_cross_compiling.sh' can be referenced or used for multiple compiles.
