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

	for i, l := range lines {
		count := strings.Split(l, ":")
		if len(count) == 2 {
			newcount := strings.Replace(count[1], "n", "", -1)
			newnewcount := strings.Replace(newcount, " ", "", -1)
			log.Printf("Line %d: %s, %s", i, l, newnewcount)
			counter, err := strconv.Atoi(newnewcount)
			if err != nil {
				cresult = cresult + counter
				log.Printf("Total seconds %d %d", cresult, counter)
			} else {
				log.Println(err.Error())
			}
		}
	}

}
