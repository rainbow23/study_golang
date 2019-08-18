package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`
}

// func (p Person) MarshalJSON() ([]byte, error) {
//     v, err := json.Marshal(&struct {
//         Name      string
//         Age       int
//         Nicknames []string
//     }{
//         Name:      "Mr." + p.Name,
//         Age:       p.Age,
//         Nicknames: p.Nicknames,
//     })
//     return v, err
/*
 * json.Marshal(p):  {"Name":"Mr.mike","Age":20,"Nicknames":["a","b","c"]}
 */
// }

func main() {
	b := []byte(`{"name":"mike", "age":20, "nicknames":["a", "b", "c"]}`)
	var person Person
	// Unmarshal 無秩序な状態
	if err := json.Unmarshal(b, &person); err != nil {
		fmt.Println(err)
	}
	fmt.Println("json.Unmarshal: ", person.Name, person.Age, person.Nicknames)
	/*
	 * json.Unmarshal:  mike 20 [a b c]
	 */

	// Marshal 整理する
	v, _ := json.Marshal(person)
	fmt.Println("json.Marshal(p): ", string(v))
	/*
	 * json.Marshal(p):  {"name":"mike","age":20,"nicknames":["a","b","c"]}
	 */
}
