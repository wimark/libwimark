#!/bin/sh

rm -f /tmp/tempdata
echo "package libwimark" >> /tmp/tempdata
echo 'import ("encoding/json")' >> /tmp/tempdata
echo 'import ("errors")' >> /tmp/tempdata
quick-enum-generator < ./enums.json >> /tmp/tempdata
gofmt "/tmp/tempdata" > ./enums.go
