# Gen

[![ci](https://github.com/sixwaaaay/gen/actions/workflows/ci.yaml/badge.svg)](https://github.com/sixwaaaay/gen/actions/workflows/ci.yaml)

Gen is a code generation tool.
The purpose of Gen is to quickly generate project code and avoid repetitive work such as writing boilerplate code.
This allows developers to focus more on implementing business logic.
Gen is developed based on the cli tool of go-zero.

The features of Gen include:

- rpc service code generation
- [WIP] kafka pub/sub code generation

## Installation

```bash
go install github.com/sixaaaay/gen@latest
```

## Usage

### Generate rpc service code

```bash
gen rpc protoc {{proto file}}  --rpc_out=. --go-grpc_out=. --go_out=.
```

## Acknowledgments

- Thanks to the development team of `go-zero` for providing protobuf parsing tools and other tools.
