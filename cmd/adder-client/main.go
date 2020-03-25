package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Dr-Evans/adder-service/rpc/adder"
)

func main() {
	client := adder.NewAdderProtobufClient("http://localhost:8080", &http.Client{})

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Adder Client")
	fmt.Print(">> What is the first number: ")

	aString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	aString = strings.Replace(aString, "\n", "", -1)

	fmt.Print(">> What is the second number: ")

	bString, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	bString = strings.Replace(bString, "\n", "", -1)

	a, err := strconv.Atoi(aString)
	if err != nil {
		panic(err)
	}

	b, err := strconv.Atoi(bString)
	if err != nil {
		panic(err)
	}

	resp, err := client.AddTwo(context.Background(), &adder.AddTwoReq{
		A: int64(a),
		B: int64(b),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sum: %d\n", resp.Sum)
}
