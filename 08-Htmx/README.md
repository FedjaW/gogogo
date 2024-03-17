# Troubleshooting

For mac users:

in `.air.toml` use:

bin = "./tmp/main"
cmd = "go build -o ./tmp/main cmd/main.go"

For windows users:

bin = ".\tmp\main.exe"
cmd = "go build -o '.\tmp\main.exe' 'cmd\main.go'"

see also: https://github.com/cosmtrek/air/issues/511#issuecomment-1922022473

## env

needed https://github.com/cosmtrek/air :

`go install github.com/cosmtrek/air@latest`

basis for all:

git clone git@github.com:ThePrimeagen/fem-htmx-proj.git

## run air

on my mac run:

`/Users/fedjabilms/go/bin/air`

open localhost:42069

## run without air

`go run cmd/main.go`