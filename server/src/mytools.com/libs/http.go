package libs

import (
	"net/http"
	"io/ioutil"
	"strings"
)

func MyGet( url string ) (string,error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "",err
	}
	defer res.Body.Close()
	return string(b),nil
}

func MyPost( url string, param string ) (string,error) {

	res, err := http.Post(url,
		"application/x-www-form-urlencoded",
			strings.NewReader(param))
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "",err
	}
	defer res.Body.Close()
	return string(b),nil

}
