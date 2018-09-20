/*
@Time : 2018/9/14 18:34 
@Author : Eather Too
@File : CallCurl
*/

package CallCurl

import (
	"APITool/decoder"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Curl struct {
	urls map[string] string //API链接
}

func New(file string) *Curl{
	curl := Curl{
		decoder.New(file).GetMap(),
	}
	return &curl
}

func (curl Curl) Get(APIMark string, data map[string] string) map[string] interface{}{
	response,_ := http.Post(APIMark, "application/x-www-form-urlencoded",strings.NewReader("h=p"))
	defer response.Body.Close()
	body,_ := ioutil.ReadAll(response.Body)
	var f interface{}
	err := json.Unmarshal(body, &f)
	if err != nil {
		fmt.Println(err)
	}
	m := f.(map[string] interface{})

	return m
}