package main

import (
	"fmt"
	"net/http"
)

func main() {
	printBanner()
	fmt.Println("-> running on port :8080")
	http.HandleFunc("/", HelloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func printBanner() {
	banner := fmt.Sprintf(
		"\n                    ____  _      __                                       " +
			"\n   ____ ___  __  __/ / /_(_)____/ /_____ _____ ____        ___  _  ______ " +
			"\n  / __ `__ \\/ / / / / __/ / ___/ __/ __ `/ __ `/ _ \\______/ _ \\| |/_/ __ \\" +
			"\n / / / / / / /_/ / / /_/ (__  ) /_/ /_/ / /_/ /  __/_____/  __/>  </ /_/ /" +
			"\n/_/ /_/ /_/\\__,_/_/\\__/_/____/\\__/\\__,_/\\__, /\\___/      \\___/_/|_/ .___/" +
			"\n                                       /____/                    /_/      " +
			"\n")
	fmt.Println(banner)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello world from multistage example!")
}
