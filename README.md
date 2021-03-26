What is this
------------

This is a simple application that computes the SHA256 of a given file using an API.

Example
-------

```
$ ./tl-hasher &
[1] 4255
2021/03/26 15:48:51 Listening on port 8080
2021/03/26 15:48:51 Application Started
$ curl -F "file=@/<myfile>" http://localhost:8080/tl-hasher/sha256
{"hash":"a6efffb83fb5469196b8cf0091cf80aefbe5b79b9ffefd0819c46be531636aba","time_taken_ms":2}
```

Build it
----

Requires Golang 1.14

```
$ go build -o tl-hasher
```