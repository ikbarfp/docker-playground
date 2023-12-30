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
		"\n                                        "+
		"\n  ____ __________ _      ___  _  ______ "+
		"\n / __ `/ ___/ __ `/_____/ _ \\| |/_/ __ \\"+
		"\n/ /_/ / /  / /_/ /_____/  __/>  </ /_/ /"+
		"\n\\__,_/_/   \\__, /      \\___/_/|_/ .___/ %v"+
		"\n          /____/               /_/      "+
		"\n", version)
	fmt.Println(banner)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello world from arg example!")
}
