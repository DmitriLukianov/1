package main

import (
	"fmt"
	"study/feature1"
	"study/feature_postgres/simple_connetion"
)

func main() {
	fmt.Println("main")
	feature1.Cikls()
	simple_connetion.CheckConnection()

}
