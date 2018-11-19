package main

import (
	"fmt"

	tc "bitbucket.org/wimarksystems/libwimark/libtc"
)

const classes_file = "../classes.yaml"

func main() {
	if err := tc.ApiTest(classes_file); err != nil {
		fmt.Println(err.Error())
	}
}
