package main

import(
	"fmt"
	"net/http"
	"io"
	"os"
)

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts") 
	if err != nil {
		fmt.Print(nil)
		return
	}
	defer resp.Body.Close() 
   	io.Copy(os.Stdout, resp.Body)
}