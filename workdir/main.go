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
		"\n                      __       ___                          "+
		"\n _      ______  _____/ /______/ (_)____      ___  _  ______ "+
		"\n| | /| / / __ \\/ ___/ //_/ __  / / ___/_____/ _ \\| |/_/ __ \\"+
		"\n| |/ |/ / /_/ / /  / ,< / /_/ / / /  /_____/  __/>  </ /_/ /"+
		"\n|__/|__/\\____/_/  /_/|_|\\__,_/_/_/         \\___/_/|_/ .___/ %v "+
		"\n                                                   /_/      "+
		"\n", version)
	fmt.Println(banner)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello world from workdir example!")
}
