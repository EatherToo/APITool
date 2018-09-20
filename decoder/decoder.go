/*
@Time : 2018/9/14 17:20 
@Author : Eather Too
@File : Decoder
*/

package decoder


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//api配置节点结构体
type node struct {
	Url string
	APIMark string
	Introduce string
}

//配置文件结构体
type config struct {
	Prefix string
	Nodes []node
}

//配置文件构造函数
func New(file string) *config {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())

	}
	conf := config{}
	err2 := json.Unmarshal(data, &conf)
	if err2 != nil {
		fmt.Println(err.Error())
	}
	return &conf
}

//将api节点转为map储存
func (conf config) GetMap() map[string] string {
	maps := make(map[string] string)
	maps["prefix"] = conf.Prefix
	for _, _node := range conf.Nodes{
		maps[_node.APIMark] = _node.Url
	}
	return maps
}