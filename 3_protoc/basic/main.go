package main

import (
	"log"

	basic "github.com/KrishnaIyer/goexamples/3_protoc/basic/gen"
)

func main() {
	req := basic.SearchRequest{
		Query: "test",
	}
	log.Print(req)
}
