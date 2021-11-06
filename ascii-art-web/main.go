package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./templates")))
	http.HandleFunc("/ascii-art", AsciiHandler)
	http.ListenAndServe(":3000", nil)
}

//handes /ascii-art request
func AsciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		banner := r.FormValue("banner")
		text := r.FormValue("text")

		for _, v := range text {
			if v != 13 && v != 10 {
				if v < 32 || v > 126 {
					http.Error(w, "500: internal server error", http.StatusInternalServerError)
					return
				}
			}

		}
		result, err := Ascii(text, banner)

		if err != nil {
			http.Error(w, "404: bannerfile not found", http.StatusNotFound)
			return
		}

		//Prints out the result of the file
		for _, word := range result {
			for _, line := range word {
				fmt.Fprintf(w, line)
				fmt.Fprintf(w, "\n")
			}
			if len(word) == 0 {
				fmt.Fprintf(w, "\n")
			}
		}

	} else {
		http.Error(w, "400: bad request", http.StatusBadRequest)
		return
	}
}

/*takes a text string and banner string and returns a two-dimensional array
holding the ascii art representations of the named banner for each word*/
func Ascii(text string, banner string) ([][]string, error) {

	var err error
	err = nil

	bannerTxt := banner + ".txt"
	var result [][]string

	bannerFile, err1 := os.ReadFile(bannerTxt)
	bannerFileSlice := strings.Split(string(bannerFile), "\n")
	textSlice := TextToArray(text)

	if err1 != nil {
		err = errors.New("missing bannerfile")
		return result, err
	}

	for _, v := range textSlice {
		lineStart := LineStartArray(v)
		if len(v) == 0 {
			result = append(result, []string{})
		} else {
			result = append(result, PrintAscii(lineStart, bannerFileSlice))
		}
	}

	return result, err

}

//Makes an array of ints that mark the starting lines of characters
func LineStartArray(s string) []int {

	var lineNumbers []int

	for i := 0; i < len(s); i++ {
		lineNr := 9 * (int(s[i]) - 32)
		lineNumbers = append(lineNumbers, lineNr)
	}

	return lineNumbers
}

//Prints out required characters
func PrintAscii(lines []int, charFile []string) []string {

	var result []string
	for i := 1; i <= 8; i++ {
		var line string
		for j := 0; j < len(lines); j++ {
			line += charFile[lines[j]+i]
		}
		result = append(result, line)

	}
	return result

}

//Takes an input string text and splits it at linebreaks
func TextToArray(text string) []string {
	var textSlice []string

	var tempWord string
	for i := 0; i < len(text); i++ {
		if text[i] == 10 {
			tempWord = ""

		} else if text[i] != 13 {
			tempWord += string(text[i])
			if i == len(text)-1 {
				textSlice = append(textSlice, tempWord)
			}
		} else if text[i] == 13 {
			textSlice = append(textSlice, tempWord)
		}
	}

	return textSlice
}
