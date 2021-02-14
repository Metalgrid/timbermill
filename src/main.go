package main

import "fmt"

func main() {
	var config *TMConfig
	config = LoadConfig("../config/timbermill.yaml")

	for _, collector := range config.Collectors {
		fmt.Printf("%+v\n", collector)
	}
}
