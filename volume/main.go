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
		"\n              __                                         "+
		"\n _   ______  / /_  ______ ___  ___        ___  _  ______ "+
		"\n| | / / __ \\/ / / / / __ `__ \\/ _ \\______/ _ \\| |/_/ __ \\"+
		"\n| |/ / /_/ / / /_/ / / / / / /  __/_____/  __/>  </ /_/ /"+
		"\n|___/\\____/_/\\__,_/_/ /_/ /_/\\___/      \\___/_/|_/ .___/ %v"+
		"\n                                                /_/      "+
		"\n", version)
	fmt.Println(banner)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	pathParam := r.URL.Path[1:]
	_, _ = fmt.Fprintf(w, "Hello %v, Cuk!", pathParam)

	dataByte := []byte("Hello " + pathParam + " , Cuk!")
	destination := os.Getenv("TEMP_DIR")

	file := destination + "/" + pathParam + ".txt"
	if err := os.WriteFile(file, dataByte, 0666); err != nil {
		_, _ = fmt.Fprintf(w, "[ERROR] failed to write file : %v", err)
	}

	fmt.Println("[INFO] writing to " + file + " succeeded")
}
