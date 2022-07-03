package main

import (
	// "context"
	"bytes"
	"fmt"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewShell("localhost:5001")

	cid, err := sh.Add(strings.NewReader("hello world!"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("added %s\n", cid)

	out := fmt.Sprintf("%s.txt", cid)
	err = sh.Get(cid, out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	data, err := sh.Cat(cid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(data)
	newStr := buf.String()
	fmt.Printf("data %s", newStr)
}
