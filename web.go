package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct{}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	respuesta, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	e := escritorWeb{}
	io.Copy(e, respuesta.Body)
}
