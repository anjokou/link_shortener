# Running

The run command is to be run from the "src" directory

```sh
go run .
```

# API

The API for this application has an HTTP API meant for public access, and an RPC API meant for inter-service communication

## HTTP

The HTTP API is for exposing the current state of links

### GET /api/links/:linkId
Returns the analytics of the link ID

An example response would be:

```json
{
	"linkId": "ABCDEFG",
	"created": 803561797000,
	"timesAccessed": 5
}
```

created is formatted as a UNIX timestamp

# RPC

## MetricsRpcActions.AddLink

This function creates a link in the link database, currently in memory

There are two parameters for this function

- LinkId           string
- CreatedTimestamp int64

This function currently retrns an int, which is unused

## MetricsRpcActions.AddToAccessedCount

This function adds the specified amount to the link's access count

There are two parameters for this function

- LinkId    string
- AddAmount uint64

This function currently retrns an int, which is unused

# Testing

To run the tests, you can use this command from the "src" directory

```sh
go test -v ./...
```

Mocks for testing generated with [mockgen](https://github.com/golang/mock)

Mocks go in src/mocks, and belong to the "mocks" package.
