package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	/* if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	} */

	response, err := http.Get("http://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal(err)

	} else {
		defer response.Body.Close()
		_, err := io.Copy(os.Stdout, response.Body)
		if err != nil {
			log.Fatal(err)
		}
	}
}