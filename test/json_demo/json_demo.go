package json_demo

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string `json:"name"`
	ServerIp   string `json:"ip"`
	ServerPort string `json:"port"`
}

func SerializeStruct() {
	server := new(Server)
	server.ServerName = "asd"
	server.ServerIp = "127.0.0.1"
	server.ServerPort = "808"

	b, err := json.Marshal(server)
	if err != nil {
		fmt.Println("marshal err :", err.Error())
	}

	fmt.Println("marshal json:", string(b))
}

func SerializeMap() {
	server := make(map[string]interface{})
	server["ServerName"] = "asd"
	server["ServerIp"] = "127.0.0.1"
	server["ServerPort"] = "8080"

	b, err := json.Marshal(server)
	if err != nil {
		fmt.Println("marshal err :", err.Error())
	}

	fmt.Println("marshal json:", string(b))
}

func DeSerializeStruct(){
	jsonString :=`{"ip":"127.0.0.1","name":"asd","port":"8080"}`
	server := new(Server)
	err := json.Unmarshal([]byte(jsonString), &server)

	if err != nil {
		fmt.Println("Unmarshal err :", err.Error())
	}

	fmt.Println("Unmarshal struct:", server)
}

func DeSerializeMap(){
	jsonString :=`{"ip":"127.0.0.1","name":"asd","port":"8080"}`
	server := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &server)

	if err != nil {
		fmt.Println("Unmarshal err :", err.Error())
	}

	fmt.Println("Unmarshal map:", server)
}
