package easy_docker

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestDockerPs(t *testing.T) {
	ps, err := Ps()
	if err != nil {
		log.Fatalln(err)
		return
	}

	marshal, err := json.Marshal(ps)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(string(marshal))
}

func TestDockerImg(t *testing.T) {
	ps, err := Images()
	if err != nil {
		log.Fatalln(err)
		return
	}

	marshal, err := json.Marshal(ps)
	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(string(marshal))
}

func TestDockerRst(t *testing.T) {
	err := Restart("ethgrey1")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
