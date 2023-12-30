package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	printBanner()
	fmt.Println("-> running on port :", port)
	http.HandleFunc("/", HelloHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func printBanner() {
	version := os.Getenv("APP_VERSION")
	banner := fmt.Sprintf(""+
		"\n    __         ____                         __    __"+
		"\n   / /_  ___  / / /___ _      ______  _____/ /___/ /"+
		"\n  / __ \\/ _ \\/ / / __ \\ | /| / / __ \\/ ___/ / __  / "+
		"\n / / / /  __/ / / /_/ / |/ |/ / /_/ / /  / / /_/ /  "+
		"\n/_/ /_/\\___/_/_/\\____/|__/|__/\\____/_/  /_/\\__,_/ %v  "+
		"\n                                                    \n", version)
	fmt.Println(banner)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello world!")
}
