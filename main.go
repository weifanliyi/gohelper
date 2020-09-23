package main

import (
	"fmt"
	"gohelper/docm"
)

func main(){
	path := "./doc"

	flag, msg := docm.PathExists(path)
	if flag {
		fmt.Println(msg)
	}

	fmt.Println(msg+"sdfsdfsd")

}