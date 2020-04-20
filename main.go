package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	b := []byte(`{"organization":{"org_id":1,"name":"Name","descr":"Descr","email":"mail@yandex.ru","phone":"+79161234567","loc":"Loc"}}`)
	b, _ = prettyPrint(b)
	fmt.Printf("%s", b)
}

/**
true
1
2.34
"gopher"
["apple","peach","pear"]
{"apple":5,"lettuce":7}
{
  "organization": {
    "org_id": 1,
    "name": "Name",
    "descr": "Descr",
    "email": "mail@yandex.ru",
    "phone": "+79161234567",
    "loc": "Loc"
  }
}
*/
