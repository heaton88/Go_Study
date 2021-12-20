package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func printBody(resq *http.Response) {
	content, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func post() {
	resq, err := http.Post("https://k8stest.golowo.com/g3-authcenter-web/oauth/token",
		"application/x-www-form-urlencoded", nil)

	if err != nil {
		panic(err)
	}
	defer func() { _ = resq.Body.Close() }()
	content, err := ioutil.ReadAll(resq.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
}

func main() {
	post()
}

/*
func requestByParams() {
	request := http.NewRequest(http.MethodPost, "https://k8stest.golowo.com/g3-authcenter-web/oauth/token")
	params := make(url.Values)
	params.Add("username", "18971575840")
	params.Add("password", "Hsw198461!")
	params.Add("grant_type", "password")
	params.Add("userType", "labormng")
	request.URL.RawQuery = params.Encode()
	resq, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resq.Body.Close() }()
	fmt.Println(resq)
}
*/

/*func post() {
	resq, err := http.Post("https://k8stest.golowo.com/g3-authcenter-web/oauth/token",
		"application/x-www-form-urlencoded", nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = resp.Body }()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("%s", content)
}


func main() {
	requestByParams()
}


*/
//func main() {
//	resp, err :=
//	http.PostForm("https://k8stest.golowo.com/g3-authcenter-web/oauth/token",
//		login_data :=
//	{
//		"username":   "18971575840",
//		"password":   "Hsw198461!",
//		"grant_type": "password",
//		"user_type":  "labormng",
//	}
//	byte, _ := json.Marshal(login_data)
//	resp, _ := http.NewRequest("Post", "https://k8stest.golowo.com/g3-authcenter-web/oauth/token", "application/x-www-form-urlencoded", string.NewrReader(string(byte)))
//	resp.Header.Add("authorization", "Basic ZzNwY2M6aXh1NHJNSUw=")
//	/resp.Header.Add("content-type", "application/x-www-form-urlencoded")
//	if err != nil {
// hand error
//panic(err)
//}
//response
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Println(err.Error())
//	}
//	log.Println(string(body))
//}

/*func main() {
  	data := login_info{
  		"username": {"18971575840"}, "pasword": {"Hsw98461!"},
  		"grant_type": {"password"}, "user_type": {"labormng"},
  	}
  	data, err := json.Marshaler(&data)
  	reader := bytes.NewReader(data)
  	resp, err := http.NewRequest("POST", "https://k8stest.golowo.com/g3-authcenter-web/oauth/token", reader)
  	resp.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  	client := &http.Client{}
  	response, err := client.Do(resp)
  	defer response.Body.Close()
  	body, err := ioutil.ReadAll(resp.Body)

  	var dataReq login_respReq
  	err = json.Unmarshal(body, &dataReq)
  	fmt.Println(json.MarshalIndent(dataReq))
  }
*/
