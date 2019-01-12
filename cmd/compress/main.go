package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		var buf bytes.Buffer
		zw := gzip.NewWriter(&buf)

		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if _, err := zw.Write([]byte(stdin.Text())); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("http://localhost:8080/?q=%s\n", buf.String())
	}
}
