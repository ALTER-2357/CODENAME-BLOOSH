package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	_ "embed"

	"github.com/gorilla/mux"
)

//go:embed "License.txt"
var s string

func main() {

	print(s, "\n", "\n", "\n")

	time.Sleep(10 * time.Second)

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		router := mux.NewRouter()

		router.HandleFunc("/upload", uploadFile).Methods("POST")

		log.Fatal(http.ListenAndServe(":3031", router))

		log.Println("backends up")
		wg.Done()
	}()

	go func() {
		fs := http.FileServer(http.Dir("./site"))
		http.Handle("/", fs)

		log.Println("forntends up, go to localhost:3030 ")
		err := http.ListenAndServe(":3030", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("backends up")

	wg.Wait()

}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	tempFile, err := ioutil.TempFile("cvs", "*.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "..........................................................................\n..........................................................................\n..........................................................................\n.........########..##........#######...#######...######..##.....##........\n.........##.....##.##.......##.....##.##.....##.##....##.##.....##........\n.........##.....##.##.......##.....##.##.....##.##.......##.....##........\n.........########..##.......##.....##.##.....##..######..#########........\n.........##.....##.##.......##.....##.##.....##.......##.##.....##........\n.........##.....##.##.......##.....##.##.....##.##....##.##.....##........\n.........########..########..#######...#######...######..##.....##........\n..........................................................................\n..........................................................................\n..........................................................................\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "we got it!!!!!!!!")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprint(w, "───────────────────█████\n───────────────────██████\n───────────────────███████\n──────────────────████████\n──────────────────████████\n─────────────────█████████\n────────────────█████████\n───────────────█████████\n──────────────█████████\n──────────────██████████████████\n────────────█████████████████████\n───────────███████████████████████\n████████─██████████████████████████\n████████─███████████████████████████\n████████─████████████████████████████\n████████─████████████████████████████\n████████─████████████████████████████\n████████─████████████████████████████\n████████─███████████████████████████\n████████─██████████████████████████\n████████─█████████████████████████\n████████─████████████████████████\n████████─███████████████████████\n")
	fmt.Fprintf(w, "\n")

}
