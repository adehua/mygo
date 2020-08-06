package xml_demo

import (
	"encoding/xml"
	"fmt"
)

type person struct {
	Name string `xml:"s,attr`
	Age  int
}

func Xml_test() {
	p := person{Name: "dd", Age: 18}
	data, err := xml.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	p2 := new(person)
	xml.Unmarshal(data, p2)

	fmt.Println(p2)
}
