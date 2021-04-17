package main

import (
	"fmt"
	"github.com/psyonly/FishersInSDV/routes"
)

func main() {

	err := routes.InitRoute()
	if err != nil {
		fmt.Println(err)
	}
}
