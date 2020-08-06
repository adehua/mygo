package point_demo

import "fmt"

func TestPoint() {
	var count int = 30
	var countPoint *int
	var countPoint1 *int
	countPoint = &count

	fmt.Printf("count 变量=的地址: %x \n", &count)
	fmt.Printf("count 存储地址: %x \n", countPoint)
	fmt.Printf("count 指针指向地址的值: %d \n", *countPoint)
	fmt.Printf("count 存储地址: %x \n", countPoint1)
	fmt.Printf("count 指针指向地址的值: %d \n", *countPoint1)
}

func TestPointArr() {
	a, b := 1, 2
	pointArr := [...]*int{&a, &b}
	fmt.Println("指针数组 pointArr : ", pointArr)

	arr := [...]int{3, 4, 5}
	arrPoint := &arr
	fmt.Println("数组指针 arrPoint : ", arrPoint)
}
