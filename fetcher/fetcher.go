package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error http code:", resp.StatusCode)
		return nil, fmt.Errorf("fetcher get %s error, code: %d", resp.StatusCode)
	}
	reader := bufio.NewReader(resp.Body)
	e := getEncoding(reader)
	//
	utf8Reader := transform.NewReader(reader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

func getEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, e := r.Peek(1024)
	if e != nil {
		log.Printf("fetcher getEncoding error:%v", e)
		return unicode.UTF8
	}
	encodig, _, _ := charset.DetermineEncoding(bytes, "")
	return encodig
}
