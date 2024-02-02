package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	values := url.Values{}
	for _, pair := range os.Args[1:] {
		k, v, ok := strings.Cut(pair, "=")
		if !ok {
			fmt.Fprintf(os.Stderr, "%v: bad format, expecting 'key=value'\n", pair)
			usage()
		}
		values.Add(k, v)
	}
	must(fmt.Fprintln(os.Stdout, values.Encode()))
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %v [key=val ...]\n", os.Args[0])
	os.Exit(1)
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func must[T any](t T, err error) T {
	check(err)
	return t
}
