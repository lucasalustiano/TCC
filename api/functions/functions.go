package functions

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Func1(w http.ResponseWriter, r *http.Request) {
    url := "https://www.google.com"

    var result [][]byte

	for i := 0; i < 10; i++ {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		result = append(result, body)
	}

	w.Write(bytes.Join(result, nil))
}

func Func2(w http.ResponseWriter, r *http.Request) {
    url := "https://google.com/"
    resp, err := http.Get(url)
    if err != nil {
        w.Write([]byte(fmt.Sprintf("URL inválida!")))
    }
    defer resp.Body.Close()

    scanner := bufio.NewScanner(resp.Body)
    scanner.Split(bufio.ScanWords)

    wordCount := 0
    for scanner.Scan() {
        wordCount++
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        w.Write([]byte(fmt.Sprintf("Erro ao ler a página!")))
    }

    w.Write([]byte(fmt.Sprintf("Número de Palavras: %d\nTamanho em bytes: %d bytes", wordCount, len(body))))
}