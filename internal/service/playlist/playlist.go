package playlist

import (
	"container/list"
	"log"
	"os"

	"github.com/galazat/MusicPlayer/internal/entity"
	"github.com/wtolson/go-taglib"
)

func NewPlaylist() *list.List {
	return list.New()
}

func ServeDir(path string, playlist *list.List) {

	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Println(err)
	}

	files, err := file.Readdir(-1)
	if err != nil {
		log.Println(err)
	}

	for i := range files {
		music := entity.NewMusic()
		music.Name = files[i].Name()

		newFile, err := taglib.Read(files[i].Name())
		if err != nil {
			log.Println(err)
		}
		music.Duration = newFile.Length()

		playlist.PushBack(music)
	}

}
