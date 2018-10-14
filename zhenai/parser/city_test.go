package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserCity(t *testing.T) {
	bytes, err := ioutil.ReadFile("city_data.html")
	if err != nil {
		panic(err)
	}
	results := ParserCity(bytes)

	fmt.Println(results)

}
