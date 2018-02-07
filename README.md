# Gorgeous

Easy to use prettifier `go test`.

## Usage

`go test -v -cover ./... | gorgeous`

Voila!

**Please note:**
The `-v` flag is important for gorgeous to pick up all the relevant information.

## Usage with docker-compose

`docker-compose` prints ansi codes as well as container prefixes, this messes with `gorgeous`.

For `gorgeous` to work in this case, please fo the following:

- run `docker-compose` with the `--no-ansi` flag
- ask `gorgeous` to strip the prefixes. Ex: `gorgeous -prefix="app_1         | "`

The prefix should be **the entire string prepended by `docker-compose`**