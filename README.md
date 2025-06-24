# go-static-site-generator

This is a project to rewrite the
[static-site-generator](https://github.com/yanmoyy/static-site-generator) in Go.

## Why Go?

- Faster than Python
- Better testability
- Type safety

## How to use

Generate html files from markdown files

```bash
go run cmd/gen/main.go
```

> Default Paths are: md for `./files/content` and html for `./files/output`.

Serve html files

```bash
go run cmd/srv/main.go
```
