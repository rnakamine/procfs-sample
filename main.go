package main

import (
	"log"

	"github.com/prometheus/procfs"
)

func main() {
	fs, err := procfs.NewFS("/proc")
	if err != nil {
		log.Fatal(err)
	}

	netTCP, err := fs.NetTCP()
	for _, v := range netTCP {
		log.Println(v)
	}
}
