package main

import (
	"fmt"
	"movieReviewSystem/movieReviewSystemApi/shared/tool"
)

func main() {
	data, err := tool.NewSnowflake(1, 1)
	if err != nil {
		panic(err)
	}
	err = data.SetModOptions(10, 0)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 1000; i++ {
		fmt.Println(data.Generate())
	}
}
