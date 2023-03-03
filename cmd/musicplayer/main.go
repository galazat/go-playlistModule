package main

import (
	"fmt"

	"github.com/wtolson/go-taglib"
)

func main() {

	// newPlaylist := playlist.NewPlaylist()
	// playlist.ServeDir("./music", newPlaylist)

	// fmt.Println(newPlaylist)

	f, _ := taglib.Read("./music/AudioFile.mp3")
	fmt.Println(f)
}
