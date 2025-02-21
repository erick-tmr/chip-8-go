package screen

import (
	"errors"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var ErrLoadMedia = errors.New("screen/load_media: error loading media")

func LoadMedia() (*sdl.Surface, func(), error) {
	helloWorldImg, err := sdl.LoadBMP("./media/hello_world.bmp")
	if err != nil {
		fmt.Println("error trying to load img:", err)
		return nil, nil, ErrLoadMedia
	}

	cleanup := func() {
		fmt.Println("cleanup load media")
		helloWorldImg.Free()
	}
	return helloWorldImg, cleanup, nil
}
