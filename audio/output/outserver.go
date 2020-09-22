package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100
const seconds = 2

func main() {
	err := portaudio.Initialize()
	if err != nil {
		fmt.Println(err)
	} else {
		defer portaudio.Terminate()
	}
	buffer := make([]float32, sampleRate*seconds)

	stream, err := portaudio.OpenDefaultStream(0, 1, sampleRate, len(buffer), func(out []float32) {
		resp, err := http.Get("http://localhost:8080/audio")
		if err != nil {
			fmt.Println(err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		responseReader := bytes.NewReader(body)
		err = binary.Read(responseReader, binary.BigEndian, &buffer)
		if err != nil {
			fmt.Println(err)
		}
		copy(out, buffer)
	})
	if err != nil {
		fmt.Println(err)
	}

	err = stream.Start()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Second * 40)
	err = stream.Stop()
	if err != nil {
		fmt.Println(err)
	}
	defer stream.Close()

	if err != nil {
		fmt.Println(err)
	}

}
