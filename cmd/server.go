package main

import (
	"fmt"
	"log"

	"github.com/go-http-server/core/utils"
)

func main() {
	env, err := utils.LoadEnviromentVariables("./")
	if err != nil {
		log.Fatal("Cannot load enviroment variables: ", err)
	}

	fmt.Println(env.DB_SOURCE)
}
