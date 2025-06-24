# go-static-site-generator

This is a project to rewrite the
[static-site-generator](https://github.com/yanmoyy/static-site-generator) in Go.

## Why Go?

- Faster than Python
- Better testability
- Type safety

## How to use

Put your markdown files in `./files/content` and generate html files in
`./files/output`.

You can also put static files in `./files/static` and they will be copied to
`./files/output`.

You can see example files in `./files/`.

### Commands

Generate html files from markdown files

```bash
go run cmd/gen/main.go
```

> Default Paths are: md for `./files/content` and html for `./files/output`.

```bash
go run cmd/srv/main.go
```
