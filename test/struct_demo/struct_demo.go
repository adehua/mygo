package struct_demo

import "fmt"

type Dog struct {
	ID   int
	Name string
	Age  int
}

func TestForStruct() {
	//var dog Dog
	//dog.ID = 1
	//dog.Name = "name"
	//dog.Age = 3

	//方式2
	//dog := Dog{ID: 2,Name: "yay",Age: 33}

	// 方式3
	dog := new(Dog)
	dog.ID = 2
	dog.Name = "sss"
	dog.Age = 44


	fmt.Println("dog is:", dog)
}

//func tests(){
//	fmt.Println("asdasd")
//}
//
//type Animal struct{
//	Colour string
//}
//
//type Dog struct{
//	Animal
//	ID int
//	Name string
//	Age int
//}
//
//type Cat struct{
//	Animal
//	ID int
//	Name string
//	Age int
//}
//
//func (d * Dog) Eat() string  {
//	fmt.Println("dog -> yummy yummy!")
//	return "dog -> yummy yummy!"
//}
//
//func (d * Dog) Run() string  {
//	fmt.Println("ID : ",d.ID, "Dog is running")
//	return "Dog is running"
//}
//
//func (d * Cat) Eat() string  {
//	fmt.Println("cat -> yummy yummy!")
//	return "cat -> yummy yummy!"
//}
//
//func (d * Cat) Run() string  {
//	fmt.Println("ID : ",d.ID, "Cat is running")
//	return "Cat is running"
//}
