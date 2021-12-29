# libwimark

Library with main models, functions and bindings, using in Golang's microservices in Wimark.

## Usage

After change any *.toml files you need to regenerate *enums.go files. It can be done with [enum-generator](https://github.com/wimark/enum-generator) tool. 

```bash
# install enum-generator

go get github.com/wimark/enum-generator
go install github.com/wimark/enum-generator

# build enums
./enums.sh

```

## Copyright

Wimark Systems, 2021