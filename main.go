package main

import (
	"net/http"

	"./ocr"
)

func main() {
	http.HandleFunc("/ocr", ocr.PDFToText)
	http.HandleFunc("/show", ocr.ShowJSON)
	http.ListenAndServe(":8091", nil)
}
