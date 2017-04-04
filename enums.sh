#!/bin/sh

rm -f /tmp/tempdata
echo "package libwimark" >> /tmp/tempdata
echo 'import ("encoding/json")' >> /tmp/tempdata
echo 'import ("errors")' >> /tmp/tempdata
echo 'import ("gopkg.in/mgo.v2/bson")' >> /tmp/tempdata
quick-enum-generator -enable-bson < ./enums.toml >> /tmp/tempdata
gofmt "/tmp/tempdata" > ./enums.go
