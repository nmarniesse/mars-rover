# Mars rover

Implementation of the [Google Mars Rover challenge](https://code.google.com/archive/p/marsrovertechchallenge/) in Go.


## Requirements

- [Go](https://go.dev) or [Docker](https://www.docker.com)

If docker, run this command to create an alias:

```bash
alias go="docker run -v ./:/app -v ./GOPATH:/go -w /app golang go"
```

## Build & run

```bash
make build
./mars-rover # Depends of your OS
```

## Run dev

```bash
make run
# or
go run .
# or run a specific input file
go run . input/challenge.txt
```

## Tests

```bash
make test
# or
go test ./...
```

## TODO

- [ ] Question: check a rover can't be at the same place than another one?
