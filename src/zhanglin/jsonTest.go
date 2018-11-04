package zhanglin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func JsonTest() {

	//---------------Marshal
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	//fmt.Println(string(b[:]))

	//---------------Unmarshal
	var jsonBlob = []byte(`[
            {"ID":1,"Name":"Reds1","Colors":["Crimson","Red1","Ruby1","Maroon1"]},
            {"ID":2,"Name":"Reds2","Colors":["Crimson","Red2","Ruby2","Maroon2"]},
            {"ID":3,"Name":"Reds3","Colors":["Crimson","Red3","Ruby3","Maroon3"]}
        ]`)

	var animals []ColorGroup
	error := json.Unmarshal(jsonBlob, &animals)
	if error != nil {
		fmt.Println("error:", error)
	}

	//fmt.Printf("%+v", animals)
	fmt.Println(animals)
	for i, x := range animals {
		fmt.Println(i, x)

	}
	//---------------Indent
	dst := new(bytes.Buffer)
	json.Indent(dst, jsonBlob, "##", "**")
	fmt.Println(dst)
}
