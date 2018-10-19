package main

import (
	"go-utils/common"
	"go-utils/logic"
	// "go-utils/logic/common"
	"fmt"
	// "github.com/bitly/go-simplejson"
	// "runtime"
	"encoding/json"
	"reflect"
)

func main() {
	name := "World"
	str := logic.Hello(name)
	fmt.Println(str)

	JsonConfig1 := common.NewJsonConfig()
	jsonConfig1, _ := JsonConfig1.Load("config/config.json")
	u, _ := jsonConfig1.Map()
	fmt.Println("JJJJJJJJ1111111", u)
	JsonConfig2 := common.NewJsonConfig()
	var jsonConfig2 interface{}
	JsonConfig2.Load2("config/config.json", &jsonConfig2)
	fmt.Println("JJJJJJJJ2222222", jsonConfig2)
	JsonConfig3 := common.NewJsonConfig()
	jsonConfig3, _ := JsonConfig3.Load("config/array.json")
	w, _ := jsonConfig3.Array()
	fmt.Println("JJJJJJJJ3333333", w)
	JsonConfig4 := common.NewJsonConfig()
	var jsonConfig4 []interface{}
	JsonConfig4.Load2("config/array.json", &jsonConfig4)
	fmt.Println("JJJJJJJJ4444444", jsonConfig4)

	countryCapitalMap := make(map[string]interface{}) /*创建集合 */

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"
	fmt.Println("MAP string interface ", countryCapitalMap)

	/* struct => map */
	jsonConfigS := []byte(`{"host": "127.0.0.1", "port": 3000}`)
	jsonConfig := struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	}{}
	json.Unmarshal(jsonConfigS, &jsonConfig)
	fmt.Println("oooooo7777", jsonConfig)
	v := reflect.ValueOf(jsonConfig)
	fmt.Println("oooooo9999", v)
	t := reflect.TypeOf(jsonConfig)
	fmt.Println("oooooo8888", t)
	var p = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("ffffffffff", t.Field(i))
		p[t.Field(i).Name] = v.Field(i).Interface()
	}
	fmt.Println("oooooo4444", p)
	res, _ := common.InvokeService("http://127.0.0.1:3001/test", "POST", p)

	type PostResp struct {
		Code int
		Data interface{}
	}
	result := PostResp{}
	json.Unmarshal(res, &result)
	fmt.Println("00000", result)

	/*
		maxCPU := runtime.NumCPU()
		runtime.GOMAXPROCS(maxCPU)
		strJson := `{"announcer": {"nickname": "非议讲史", "kind": "user", "updated_at": 1494983507000, "track_id": 38088960}}`
		mapJson, _ := simplejson.NewJson([]byte(strJson))
		println(mapJson)
		fmt.Println(maxCPU)
		fmt.Println(mapJson)
		//  beego.Run()
	*/

}
