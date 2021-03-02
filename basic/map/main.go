package main

import "fmt"

func main() {

	string_map := map[string]string{}
	string_map["name"] = "programming sesang"
	string_map["firstname"] = "firstname"
	string_map["secondname"] = "secondname"
	string_map["thirdname"] = "thirdname"

	fmt.Println("before result : ", string_map)

	delete(string_map, "firstname")
	delete(string_map, "secondname")
	delete(string_map, "thirdname")

	fmt.Println("after result : ", string_map)
}
