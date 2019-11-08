package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file := "/etc/nginx/nginx.conf"
	port := os.Getenv("PORT")

	cont, err := ioutil.ReadFile(file)

	if err != nil {
		panic(err)
	}

	contStr := string(cont)

	contStr = strings.Replace(contStr, "$PORT", port, 1)

	err = ioutil.WriteFile(file, []byte(contStr), 0644)

	if err != nil {
		panic(err)
	}

	fmt.Print(contStr)
}
