# Mars rover

Implementation of the [Google Mars Rover challenge](https://code.google.com/archive/p/marsrovertechchallenge/) in Go.


## Requirements

- [Go](https://go.dev) or [Docker](https://www.docker.com)

If docker, run this command to create an alias:

```bash
alias go="docker run -v ./:/app -v ./GOPATH:/go -w /app golang go"
```

## Run

```bash
go run .
go run . input/challenge.txt
```

## Tests

```bash
go test ./...
```

## TODO

- [ ] Question: check a rover can't be at the same place than another one?
