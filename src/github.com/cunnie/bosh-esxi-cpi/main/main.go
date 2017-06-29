package main

import (
	"fmt"
	"flag"
)

var (
	configPathOpt = flag.String("configPath", "", "Path to configuration file")
)

func main() {
	flag.Parse()
	
	fmt.Println("I love muh dawg!")
}
