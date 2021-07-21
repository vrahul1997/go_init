package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"example.com/proto_buff_go/src/enum_example"
	"example.com/proto_buff_go/src/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()

	// function for read and write combined
	readAndWrite("simple.txt", sm)
	smAsString := toJson(sm)
	fmt.Println(smAsString)

	sm3 := &simple.SimpleMessage{}
	fromJSON([]byte(smAsString), sm3)
	fmt.Println(sm3.Name)
}

// function for protobuff to json
func toJson(pb proto.Message) string {
	output, err := protojson.Marshal(pb)
	if err != nil {
		log.Fatal("error marshaling to json", err)
		return ""
	}
	return string(output)
}

// function for from JSON
func fromJSON(in []byte, pb proto.Message) {
	err := protojson.Unmarshal(in, pb)
	if err != nil {
		log.Fatal("Couldnt unmarshal the json", err)
	}
}

// functio for reading and writing at the same time
func readAndWrite(filename string, pb proto.Message) {
	writeToFile(filename, pb)
	sm2 := &simple.SimpleMessage{}
	readFromFile(filename, sm2)
	fmt.Println(sm2)
}

// write the message to file
func writeToFile(filename string, pb proto.Message) error {
	output, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Cannot serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(filename, output, 0644); err != nil {
		log.Fatal("Cannot write to file")
		return err
	}
	return nil
}

// read from file
func readFromFile(filename string, pb proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cannot deserialize the file", err)
		return err
	}
	unmarshall_err := proto.Unmarshal(data, pb)
	if unmarshall_err != nil {
		return unmarshall_err
	}
	return nil
}

// simple function for getting the message from protofile
func doSimple() *simple.SimpleMessage {
	sm := simple.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "John Rahull",
		SampleList: []int32{1, 2, 3, 4, 5},
	}

	return &sm
}

func doENUM() *enum_example.EnumMessage {
	ep := enum_example.EnumMessage{
		Id:           2002,
		DayOfTheWeek: 6,
	}
	return &ep

}
