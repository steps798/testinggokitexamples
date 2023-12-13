# Examples

This repository contains examples extracted from the core [kit](https://github.com/go-kit/kit) repository.

For more information about these examples,
 including a walkthrough of the stringsvc example,
 see [gokit.io/examples](https://gokit.io/examples).

To skip possibly time-consuming tests, use this command.
```shell
$ go test -json ./... -short
```

To generate coverage report of an example, e.g. /shipping, use this command.
```shell
$ go test -coverpkg=./shipping/... -coverprofile=coverage.out ./shipping/...
$ go tool cover -html=coverage.out
```
