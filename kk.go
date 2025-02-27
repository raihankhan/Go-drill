package main

import (
	"fmt"
	"time"
)

type File interface{}

type Audio struct {
	File
	Duration time.Duration
}

type Image struct {
	File
	Width  uint
	Height uint
}

func main() {
	var files = map[string]File{
		"1": Audio{
			Duration: 14 * time.Second,
		},
		"2": Image{
			Height: 9989,
			Width:  1111,
		},
		"3": Image{
			Width:  1234,
			Height: 5678,
		},
	}
	for k, v := range files {
		switch file := v.(type) {
		case Audio:
			fmt.Printf("%s: Audio %d seconds", k, file.Duration/time.Second)
		case Image:
			fmt.Printf("%s: Image %dx%d", k, file.Width, file.Height)
		}
	}
}
