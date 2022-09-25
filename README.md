# krampus
Krampus is a barebone replacement for Python's http.server, but with uploading, written in Go.

### Flags
* --ssl true (default false)
* --ssl port (default 9001)

### SSL/TLS
```
$ ls
cert.pem  key.pem
$ doas ./krampus --ssl true --port 8443
Server starting at port: 8443
```

## Uploading
### Windows
C:\\> curl -F "file=@\<file.txt\>" http[:]//127.0.0.1:9001/upload

### Unix
$ curl -F "file=@\<file.txt\>" http[:]//127.0.0.1:9001/upload