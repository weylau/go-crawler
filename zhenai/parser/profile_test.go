package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserProfile(t *testing.T) {
	bytes, err := ioutil.ReadFile("profile_data.html")
	if err != nil {
		panic(err)
	}
	results := ParserProfile(bytes, "test")

	fmt.Println(results)

}
