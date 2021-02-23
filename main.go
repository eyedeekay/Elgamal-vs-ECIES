package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	file := flag.String("file", "ecies.idk.md", "filename to read")
	flag.Parse()
	bytes, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	cresult := 0
	lines := strings.Split(string(bytes), "\n")

	for _, l := range lines {
		count := strings.Split(l, ":")
		if len(count) == 2 {
			newcount := strings.Trim(count[1], " n\t\n")
			counter, err := strconv.Atoi(string(newcount))
			if err == nil {
				cresult = cresult + counter

			} else {
				log.Println(err.Error())
			}
		}
	}
	log.Printf("Total Nanoseconds %d", cresult)
}
