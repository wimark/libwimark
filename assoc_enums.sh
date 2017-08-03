#!/bin/sh

rm -f /tmp/tempdata
echo "package libwimark" >> /tmp/tempdata
assoc-enum-generator < ./assoc_enums.toml >> /tmp/tempdata
gofmt "/tmp/tempdata" > ./assoc_enums.go
