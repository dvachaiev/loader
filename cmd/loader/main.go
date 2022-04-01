package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"plugin"
	"time"
)

const (
	path       = "./plugin.so"
	methodName = "Run"
)

func main() {
	url := os.Args[1]

	if err := download(url, path); err != nil {
		panic(err)
	}

	pl, err := plugin.Open(path)
	if err != nil {
		panic(fmt.Errorf("load plugin: %w", err))
	}

	symb, err := pl.Lookup(methodName)
	if err != nil {
		panic(fmt.Errorf("entry point lookup: %w", err))
	}

	method, ok := symb.(func())
	if !ok {
		panic(errors.New("entry point: wrong signature"))
	}

	method()
}

func download(url string, path string) error {
	cl := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := cl.Get(url)
	if err != nil {
		return fmt.Errorf("making request: %w", err)
	}

	defer resp.Body.Close()

	fobj, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	defer fobj.Close()

	if _, err = io.Copy(fobj, resp.Body); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}
