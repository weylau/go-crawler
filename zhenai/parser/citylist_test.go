package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

const ResultSize = 470

func TestParserCityList(t *testing.T) {
	bytes, err := ioutil.ReadFile("citylist_data.html")
	if err != nil {
		panic(err)
	}
	results := ParserCityList(bytes)

	expectedUrl := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedItem := []string{
		"City:阿坝", "City:阿克苏", "City:阿拉善盟",
	}

	for i, url := range expectedUrl {
		if results.Results[i].Url != url {
			t.Errorf("expected url %s,but was #%d:%s;", url, i, results.Results[i].Url)
		}
	}

	for i, item := range expectedItem {
		if results.Items[i].(string) != item {
			t.Errorf("expected item %s,but was #%d:%s;", item, i, results.Items[i])
		}
	}

	if len(results.Results) != ResultSize {
		t.Errorf("test result lenth should have %d, the result lenth is %d", ResultSize, len(results.Results))
	} else {
		fmt.Println("test ok!!!")
	}

}
