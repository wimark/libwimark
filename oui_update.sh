#!/bin/bash

URL="https://linuxnet.ca/ieee/oui/nmap-mac-prefixes"

TEMP_FILE=/tmp/nmap_oui.txt
TEMP_GOFILE=/tmp/oui_temp.go 
OUI_FILE=oui.go
wget -O $TEMP_FILE $URL -o /dev/null

echo 'package libwimark' > $TEMP_GOFILE
echo "" >> $TEMP_GOFILE
echo 'var ManufacturerMap = map[string]string{' >> $TEMP_GOFILE

## some fucking cycle
awk -F"\t" '!_[$1]++' $TEMP_FILE >> $TEMP_FILE.tmp 
# exit 0
sed -e 's/\(.*\)\t\(.*\)/\t"\L\1": "\E\2",/' $TEMP_FILE.tmp >> $TEMP_GOFILE


echo  '}' >> $TEMP_GOFILE

rm $TEMP_FILE.tmp $TEMP_FILE
mv $TEMP_GOFILE $OUI_FILE
