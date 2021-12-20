package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func printBody(resq *http.Response) {
	content, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func resqByParams() {
	request, err := http.NewRequest(http.MethodPost, "https://k8stest.golowo.com/g3-authcenter-web/oauth/token", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("content", "application/x-www-form-urlencoded")
	request.Header.Add("authorization", "Basic ZzNwY2M6aXh1NHJNSUw=")
	params := make(url.Values)
	params.Add("username", "18971575840")
	params.Add("password", "Hsw198461!")
	params.Add("grant_type", "password")
	params.Add("userType", "labormng")
	request.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	printBody(r)
}

//jsonData := []bytese(`resqByParams()`)

func main() {
	resqByParams()
}
