package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"time"
)

func main() {
	var cepFlag string
	flag.StringVar(&cepFlag, "zipcode", "", "Argument for the zipcode to be searched.")
	flag.Parse()

	if cepFlag == "" {
		panic("A zipcode is needed to run the program. You can provide the zipcode using the -zipcode flag. Example: go run main.go -zipcode=01153000")
	}

	ctx, done := context.WithTimeout(context.Background(), 1*time.Second)
	defer done()

	viacep := NewViaCepService()
	brasilapi := NewBrasilApiService()

	c1 := make(chan *ViaCepResponse)
	c2 := make(chan *BrasilApiResponse)

	go func() {
		viaCepResponse, err := viacep.SearchZipcode(cepFlag, ctx)
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				fmt.Println("Error while searching zipcode in ViaCep:", err)
			}
		}
		c1 <- viaCepResponse
		done()
	}()

	go func() {
		brasilApiResponse, err := brasilapi.SearchZipcode(cepFlag, ctx)
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				fmt.Println("Error while searching zipcode in BrasilApi:", err)
			}
		}
		c2 <- brasilApiResponse
		done()
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Println("ViaCep response: ", msg)
			return
		case msg := <-c2:
			fmt.Println("BrasilApi response: ", msg)
			return
		case <-ctx.Done():
			fmt.Println("Timeout reached. Exiting program.")
			return
		}
	}

}
