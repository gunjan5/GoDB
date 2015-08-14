package main

import (
	"fmt"
	"net/http/httptest"
)

func main() {
	fmt.Print("")
	fmt.Print(httptest.DefaultRemoteAddr)

}
