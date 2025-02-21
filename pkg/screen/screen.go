package screen

import (
	"errors"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Sentinel errors
var ErrScreen error = errors.New("screen: error in screen package")
var ErrScreenInit error = fmt.Errorf("screen/init: error initializing screen / %v", ErrScreen)
var ErrScreenClear error = fmt.Errorf("screen/clear: error clearing screen / %v", ErrScreen)
var ErrScreenUpdate error = fmt.Errorf("screen/clear: error updating screen / %v", ErrScreen)

const SCREEN_WIDTH int32 = 640             //chip 8 x10 resolution
const SCREEN_HEIGHT int32 = 320            //chip 8 x10 resolution
const FOREGROUND_COLOR uint32 = 0x00000000 // black
const BACKGROUND_COLOR uint32 = 0xFFFFFFFF // white
const DELAY_TIME uint32 = 16               // chip 8 is 60Hz / FPS

type Screen struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func New() Screen {
	return Screen{}
}

func (s *Screen) Init() (func(), error) {
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		fmt.Println("Error while initializing SDL:", err)
		return nil, ErrScreenInit
	}

	s.window, err = sdl.CreateWindow(
		"CHIP-8 Go - by Erick Takeshi",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		SCREEN_WIDTH,
		SCREEN_HEIGHT,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		fmt.Println("Error while creating SDL Window:", err)
		return nil, ErrScreenInit
	}

	s.renderer, err = sdl.CreateRenderer(s.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Error while creating SDL Renderer:", err)
		return nil, ErrScreenInit
	}

	cleanup := func() {
		s.renderer.Destroy()
		s.window.Destroy()
		sdl.Quit()
	}

	return cleanup, nil
}

func (s *Screen) Clear() error {
	if s.renderer == nil {
		fmt.Println("You need to initialize screen before clearing it")
		return ErrScreenClear
	}

	// Background color bytes
	r := byte(BACKGROUND_COLOR >> 24)
	g := byte((BACKGROUND_COLOR >> 16) & 0xFF)
	b := byte((BACKGROUND_COLOR >> 8) & 0xFF)
	a := byte(BACKGROUND_COLOR & 0xFF)

	err := s.renderer.SetDrawColor(r, g, b, a)
	if err != nil {
		fmt.Println("Error while setting SDL Renderer draw color:", err)
		return ErrScreenClear
	}
	s.renderer.Clear()

	return nil
}

func (s *Screen) Update() error {
	if s.renderer == nil {
		fmt.Println("You need to initialize screen before updating it")
		return ErrScreenUpdate
	}

	s.renderer.Present()

	return nil
}
