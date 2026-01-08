alias b := build
alias t := test

build:
  go build .

test:
  go test ./test -v
