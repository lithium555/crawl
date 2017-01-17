package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const url = "http://www.devdungeon.com/content/working-files-go#compress" //""http://hard.rozetka.com.ua/dell_se2717h/p12469494/"

func main() {
	fmt.Println("Downloading " + url + "...")
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("File didn`t download, err = '%v'\n", err)
	}
	defer resp.Body.Close()
	// Create .gz file to write to
	file, err := os.Create("Monitor27.gz")
	if err != nil {
		log.Printf("File didn`t create, err = '%v'\n", err)
	}
	defer file.Close()

	//io.Copy(file, resp.Body)
	/*
		io.Copy is a nice little function that take a reader interface
		and writer interface, reads data from one and writes it to the other.
		Very useful for this kind of stuff!
	*/
	// Create a gzip writer on top of file writer

	gzipWriter := gzip.NewWriter(file)
	gzipWriter.Name = "BLYAD"
	defer gzipWriter.Close()
	// When we write to the gzip writer
	// it will in turn compress the contents
	// and then write it to the underlying
	// file writer as well
	// We don't have to worry about how all
	// the compression works since we just
	// use it as a simple writer interface
	// that we send bytes to
	//_, err = gzipWriter.Write([]byte("Gophers rule!\n"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Compressed data written to file.")

	path, _ := filepath.Abs(file.Name())
	fmt.Println("FILEPATH = " + path)
	log.Printf("respBody = '%v'\n", resp.Body)
	//
	//responseData,err := ioutil.ReadAll(resp.Body)
	//responseString := string(responseData)
	//fmt.Println(responseString)

	//bytes.NewReader(responseData)
	// io.Copy( gzipWriter, resp.Body)
	//	mas, err := gzipWriter.Write(responseData)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	fmt.Println("mas = '%v'\n", mas)
	//	//scanner := bufio.NewScanner(gzipWriter)
	//	_, err := io.Copy(os.Stdout, resp.Body)

	//_, err = gzipWriter.Write(responseData)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//io.Copy(gzipWriter, bytes.NewReader(resp.Body))
	io.Copy(gzipWriter, resp.Body)
}
