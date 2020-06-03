package main

import (
	"fmt"
	"goProject/pipeline"
	"os"
	"strconv"
)

func main() {
	// src % time go run main.go 2
	// 2
	// pipeline use :
	// 1 4 9 16 25
	// go run main.go 2  0.25s user 0.18s system 51% cpu 0.857 total

	// 	src % time go run main.go 1
	// 1
	// pipepine fan use :
	//   1  4  9 16 25
	// go run main.go 1  0.22s user 0.18s system 116% cpu 0.348 total

	res, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("%#v\n", res)
	if res == 1 {

		fmt.Println("pipepine fan use : ")
		pipeline.ConsumerFan()

	} else {

		fmt.Println("pipeline use : ")
		pipeline.Consumer()
	}
}
