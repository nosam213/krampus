# krampus
Krampus is a barebone replacement for Python's http.server, but with uploading, written in Go.

curl -F "file=@filename.txt" http://<server>:9001/upload
