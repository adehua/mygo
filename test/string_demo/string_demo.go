package string_demo

import (
	"fmt"
	"strconv"
	"strings"
)

func Test_string() {
	// 查找
	s := "hello world"
	fmt.Println(strings.Contains(s, "hello"), strings.Contains(s, "?"))
	fmt.Println(strings.Index(s, "ello"))

	// 合并切割
	ss := "1#2#3#4"
	// 合并
	sss1 := strings.Split(ss, "#")
	sss2 := strings.Join(sss1, "#")
	fmt.Println(sss1, sss2)

	fmt.Println(strings.HasPrefix(s, "he"), strings.HasSuffix(s, "d"))

}

func Test_string1() {
	// 基本转换
	fmt.Println(strconv.Itoa(10))
	fmt.Println(strconv.Atoi("100"))


	// 解析
	fmt.Println(strconv.ParseBool("TRUE"))
	fmt.Println(strconv.ParseFloat("3.144",32))
	fmt.Println(strconv.ParseFloat("3.146",32))

	// 格式化
	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(123,2))
}