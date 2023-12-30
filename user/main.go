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
		"\n                                            "+
		"\n  __  __________  _____      ___  _  ______ "+
		"\n / / / / ___/ _ \\/ ___/_____/ _ \\| |/_/ __ \\"+
		"\n/ /_/ (__  )  __/ /  /_____/  __/>  </ /_/ /"+
		"\n\\__,_/____/\\___/_/         \\___/_/|_/ .___/ %v"+
		"\n                                   /_/      "+
		"\n", version)
	fmt.Println(banner)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello world from user example!")
}
