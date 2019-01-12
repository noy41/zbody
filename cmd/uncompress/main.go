package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"net/url"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		u, err := url.Parse(stdin.Text())
		if err != nil {
			log.Fatal(err)
		}

		var buf bytes.Buffer
		buf.Write([]byte(u.Query()["q"][0]))

		zr, err := gzip.NewReader(&buf)
		if err != nil {
			log.Fatal(err)
		}

		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}

		if err := zr.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
