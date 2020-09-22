package main

import (
	"encoding/binary"
	"net/http"

	"github.com/gordonklaus/portaudio"
)

const sampleRate = 44100
const seconds = 2

func main() {

	err := portaudio.Initialize()
	if err != nil {
		panic(err)
	} else {
		defer portaudio.Terminate()
	}
	buffer := make([]float32, sampleRate*seconds)
	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(buffer), func(in []float32) {
		copy(buffer, in)
		//		for i := range buffer {
		//			buffer[i] = in[i]
		//		}
	})

	if err != nil {
		panic(err)
	}

	err = stream.Start()
	defer stream.Close()

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/audio", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected to be an http.Flusher")
		}

		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Content-Type", "audio")
		for {
			_ = binary.Write(w, binary.BigEndian, &buffer)
			flusher.Flush() // Trigger "chunked" encoding and send a chunk...
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
