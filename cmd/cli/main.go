package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

const (
	cmdPrint = "print"
)

func main() {
	log.Print("deky cli")
	if len(os.Args) == 1 {
		log.Fatal("expected subcommand")
	}

	flag.NewFlagSet(cmdPrint, flag.ExitOnError)

	switch os.Args[1] {
	case cmdPrint:
		query := ""
		for i, v := range os.Args[2:] {
			if i > 0 {
				query += " "
			}
			query += v
		}
		queryEsc := url.QueryEscape(query)
		fmt.Println(queryEsc)
		http.Get("http://raspberrypizero.local:8080/" + queryEsc)
	}
}
