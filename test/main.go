package main

import (
	"mygo/test/point_demo"
	"mygo/test/xml_demo"
)

func main() {

	//Xml_test()
	//     makeSlice()

	//makeMap()
	//struct_demo.TestForStruct()

	// 测试接口定义变量
	//dog := new(struct_demo.Dog)
	//cat := new(struct_demo.Cat)
	//action(dog)
	//action(cat)

	// 并发
	//fmt.Println("cpu 个数：",runtime.NumCPU())
	//runtime.GOMAXPROCS(runtime.NumCPU())
	//go gorotine.Loop()
	//go gorotine.Loop()
	//time.Sleep(time.Second * 10)

	// 协程通信
	// 启动发送数据
	//go gorotine.Send()
	// 启动接收数据
	//go gorotine.Receive()
	//time.Sleep(time.Second * 60)

	// 协程同步
	//gorotine.Read()
	//go gorotine.Write()
	//gorotine.WG.Wait()
	//fmt.Println("all done")
	//
	//time.Sleep(time.Second * 60)

	// 序列化json
	//json_demo.SerializeStruct()
	//json_demo.SerializeMap()

	// 反序列化 json
	//json_demo.DeSerializeStruct()
	//json_demo.DeSerializeMap()

	point_demo.TestPointArr()
}

//func String_test() {
//	string_demo.Test_string1()
//}

func Xml_test() {
	xml_demo.Xml_test()
}

//
//func action (b interface_demo.Behavior) string {
//	b.Run()
//	b.Eat()
//	return ""
//}

// func makeSlice(){
//     mSlice := make([]string,3)
//     mSlice[0]= 'dog'
//     mSlice[1]= 'cat'
//     mSlice[2]= 'tiger'
//     fmt.Println(mSlice)
// }

//func makeMap(){
//	mMap := make(map[int]string)
//	mMap[100] = "sss"
//	fmt.Println(mMap)
//}
