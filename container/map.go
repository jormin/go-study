package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]string)
	fmt.Println(m2)

	var m3 map[string]string
	fmt.Println(m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println(k)
	}
	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if tmpValue, ok := m["wrongkey"]; ok {
		fmt.Println(tmpValue, ok)
	} else {
		fmt.Println("Wrong key")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
