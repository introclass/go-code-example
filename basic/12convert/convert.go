package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Attr struct {
	Value string `yaml:"value"`
}

type Student struct {
	Name  string       `yaml:"name"`
	Attrs map[int]Attr `yaml:"attrs"`
}

func Base64Convert() {
	fmt.Println("###### Base64Convert #####")
	msg := "Hello, 世界"
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("base64: %s", encoded)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
}

func UintAndInt() {
	var uint_a uint64 = 100000
	fmt.Printf("uint: %d to int: %d\n", uint_a, int64(uint_a))
}

func Floor() {
	var a int32 = 80
	var b int32 = 100
	c := int32(math.Floor(float64(b) * (float64(a) / 100)))
	fmt.Println("%d", c)

	a = -3
	fmt.Println("absoulte is: ", math.Abs(float64(a)))
	fmt.Printf("current time is: %s\n", time.Now().Format("2020-02-20T03:03:07Z"))
}

func Yaml() {
	s := Student{
		Name: "lijiaocn",
		Attrs: map[int]Attr{
			0: {Value: "hello1"},
			1: {Value: "hello2"},
		},
	}

	if out, err := yaml.Marshal(s); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s\n", out)
	}

}
func Json() {
	s := Student{
		Name: "lijiaocn",
		Attrs: map[int]Attr{
			0: {Value: "hello1"},
			1: {Value: "hello2"},
		},
	}

	studentss := make([]Student, 10, 10)
	studentss = append(studentss, s, s, s, s)
	content, _ := json.Marshal(studentss)
	fmt.Printf("%s\n", string(content))

	ints := "[1,2,3]"
	intlist := make([]int, 0)
	if err := json.Unmarshal([]byte(ints), &intlist); err != nil {
		log.Panic(err)
	}
	fmt.Println(intlist)
}

func Json2() {
	s := Student{
		Name: "lijiaocn",
		Attrs: map[int]Attr{
			0: {Value: "hello1"},
			1: {Value: "hello2"},
		},
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(s); err != nil {
		log.Panic(err.Error())
	}
	fmt.Printf("%s", buf.String())
}

func main() {
	Base64Convert()
	UintAndInt()
	Json()
	Json2()
}
