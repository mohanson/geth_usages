package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	set := map[string]struct{}{}

	f, err := os.Open("eth_address.log")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		if len(line) != 42 {
			log.Fatalln(line)
		}
		set[line] = struct{}{}
	}
	f.Close()

	log.Println("Count of addresses", len(set))

	saveFile, err := os.OpenFile("eth_address_uniq.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer saveFile.Close()
	w := bufio.NewWriter(saveFile)

	for e := range set {
		w.Write([]byte(e))
		w.Write([]byte("\n"))
	}
	w.Flush()
}
