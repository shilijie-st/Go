/*
示例：使用匿名结构体分离JSON数据

*/
package main

import (
	"encoding/json"
	"fmt"
)

//Screen 定义手机屏幕
type Screen struct {
	Size       float32 //屏幕尺寸
	ResX, ResY int     //屏幕水平和垂直分辨率
}

//Battery 定义电池
type Battery struct {
	Capacity int //容量
}

//genJsonData 生成json数据
func genJSONData() []byte {
	//完整数据结构
	raw := &struct {
		Screen
		Battery
		HasTouchID bool //序列化时添加字段
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		}, Battery: Battery{
			2910,
		},
		HasTouchID: true,
	}
	//将数据序列化为JSON
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func main() {
	//生成一段JSON数据
	jsonDate := genJSONData()
	fmt.Println(jsonDate)
	//只需要屏幕和指纹识别的结构和实例
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	//反序列化到screenAndTouch
	json.Unmarshal(jsonDate, &screenAndTouch)
	//输出screenAndTouch的详细结构和实例
	fmt.Printf("%+v\n", screenAndTouch)
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	json.Unmarshal(jsonDate, &batteryAndTouch)
	fmt.Printf("%+v\n", batteryAndTouch)
}
