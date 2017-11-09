#!/bin/sh

for i in `find ./ -name "*.toml"`
do 
	cat $i | enum-generator -enable-json -enable-bson -package=libwimark | gofmt > "${i%.*}".go
done
