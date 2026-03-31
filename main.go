package main

import (
	"fmt"

	"github.com/kenvez/hostctl/internal/blocker"
	"github.com/kenvez/hostctl/internal/reader"
)

func main() {
	f, err := reader.ParseHosts()

	if err != nil {
		fmt.Println(err)
	}

	blocker.Block(f, "reddit.com")
	blocker.Save(f)
}
