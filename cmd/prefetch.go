package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"www.velocidex.com/golang/binparsergen"
	prefetch "www.velocidex.com/golang/go-prefetch"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, arg := range args {
		fd, err := os.Open(arg)
		binparsergen.FatalIfError(err, fmt.Sprintf("Open file: %v", err))

		prefetch_obj, err := prefetch.LoadPrefetch(fd)
		binparsergen.FatalIfError(err, fmt.Sprintf("Parsing Error: %v", err))

		serialized_content, _ := json.MarshalIndent(prefetch_obj, " ", " ")
		fmt.Println(string(serialized_content))
	}
}
