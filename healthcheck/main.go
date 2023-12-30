package main

import (
	"fmt"
	"net/http"
	"os"
)

var counter = 0

func main() {
	port := os.Getenv("APP_PORT")
	printBanner()
	fmt.Println("-> running on port :", port)
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/health", HealthcheckHandler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func printBanner() {
	version := os.Getenv("APP_VERSION")
	banner := fmt.Sprintf(""+
		"\n    __               ____  __         __              __                       "+
		"\n   / /_  ___  ____ _/ / /_/ /_  _____/ /_  ___  _____/ /__      ___  _  ______ "+
		"\n  / __ \\/ _ \\/ __ `/ / __/ __ \\/ ___/ __ \\/ _ \\/ ___/ //_/_____/ _ \\| |/_/ __ \\"+
		"\n / / / /  __/ /_/ / / /_/ / / / /__/ / / /  __/ /__/ ,< /_____/  __/>  </ /_/ /"+
		"\n/_/ /_/\\___/\\__,_/_/\\__/_/ /_/\\___/_/ /_/\\___/\\___/_/|_|      \\___/_/|_/ .___/  %v"+
		"\n                                                                      /_/      "+
		"\n", version)
	fmt.Println(banner)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello world from healthcheck example!")
}

func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
	counter++

	if counter > 5 {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, "I'm Sick OMG !!!")
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "I'm healthy :D")
	return
}
