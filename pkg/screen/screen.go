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

const CHIP8_SCREEN_WIDTH int32 = 64        //chip 8 resolution
const CHIP8_SCREEN_HEIGHT int32 = 32       //chip 8 resolution
const SCALE_FACTOR int32 = 10              // scale factor to make it visible
const FOREGROUND_COLOR uint32 = 0x33FF3300 // green
const BACKGROUND_COLOR uint32 = 0x00000000 // black
const DELAY_TIME uint32 = 16               // chip 8 is 60Hz / FPS or ~16.67 ms

type Screen struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	pixels   [CHIP8_SCREEN_WIDTH * CHIP8_SCREEN_HEIGHT]bool
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
		CHIP8_SCREEN_WIDTH*SCALE_FACTOR,
		CHIP8_SCREEN_HEIGHT*SCALE_FACTOR,
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
	r, g, b, a := bytesToRGBA(BACKGROUND_COLOR)

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

	rect := sdl.Rect{
		X: 0,
		Y: 0,
		W: SCALE_FACTOR,
		H: SCALE_FACTOR,
	}
	bgR, bgG, bgB, bgA := bytesToRGBA(BACKGROUND_COLOR)
	fgR, fgG, fgB, fgA := bytesToRGBA(FOREGROUND_COLOR)

	for i, pOn := range s.pixels {
		// Translate 1D index i value to 2D X/Y coordinates
		pX, pY := indexToCoordinates(int32(i))

		rect.X = pX * SCALE_FACTOR
		rect.Y = pY * SCALE_FACTOR

		if pOn {
			// pixel is on, draw foreground color
			err := s.renderer.SetDrawColor(fgR, fgG, fgB, fgA)
			if err != nil {
				fmt.Println("Could not set draw color on renderer:", err)
				return ErrScreenUpdate
			}

			err = s.renderer.FillRect(&rect)
			if err != nil {
				fmt.Println("Could not fill rect on renderer:", err)
				return ErrScreenUpdate
			}

			continue
		}

		// pixel is off, draw background color
		err := s.renderer.SetDrawColor(bgR, bgG, bgB, bgA)
		if err != nil {
			fmt.Println("Could not set draw color on renderer:", err)
			return ErrScreenUpdate
		}

		err = s.renderer.FillRect(&rect)
		if err != nil {
			fmt.Println("Could not fill rect on renderer:", err)
			return ErrScreenUpdate
		}
	}

	s.renderer.Present()

	return nil
}

func (s *Screen) SetPixelOn(x int32, y int32) {
	i := coordinatesToIndex(x, y)
	s.pixels[i] = true
}

func bytesToRGBA(bytes uint32) (byte, byte, byte, byte) {
	r := byte(bytes >> 24)
	g := byte((bytes >> 16) & 0xFF)
	b := byte((bytes >> 8) & 0xFF)
	a := byte(bytes & 0xFF)

	return r, g, b, a
}

func indexToCoordinates(i int32) (x int32, y int32) {
	// verify if i is less than 64x32 - 1 (maximum i value for screen res)

	x = i % CHIP8_SCREEN_WIDTH
	y = i / CHIP8_SCREEN_WIDTH

	return x, y
}

func coordinatesToIndex(x int32, y int32) (i int32) {
	// verify if x is less than 64 -1
	// and y is less than 32 -1 (maximum value for screen res)

	i = y*CHIP8_SCREEN_WIDTH + x

	return i
}
