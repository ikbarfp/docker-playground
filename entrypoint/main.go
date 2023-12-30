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
		"\n              __                         _       __                      "+
		"\n  ___  ____  / /________  ______  ____  (_)___  / /_      ___  _  ______ "+
		"\n / _ \\/ __ \\/ __/ ___/ / / / __ \\/ __ \\/ / __ \\/ __/_____/ _ \\| |/_/ __ \\"+
		"\n/  __/ / / / /_/ /  / /_/ / /_/ / /_/ / / / / / /_/_____/  __/>  </ /_/ /"+
		"\n\\___/_/ /_/\\__/_/   \\__, / .___/\\____/_/_/ /_/\\__/      \\___/_/|_/ .___/ %v"+
		"\n                   /____/_/                                     /_/      "+
		"\n", version)
	fmt.Println(banner)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello world from entrypoint example!")
}
