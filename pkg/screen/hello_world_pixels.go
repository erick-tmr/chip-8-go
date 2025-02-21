package screen

var HELLO_WORLD_PIXELS = [64 * 32]bool{
	false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, true, true, true, true, true, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, true, true, true, true, true, false, false, false, false, false, true, true, true, true, true, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, false, false, false, false, false,
	false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, false, false, false, true, true, true, true, true, false, false, false, false, false, true, true, true, true, true, false, false, false, false, false, true, true, true, true, true, false, false, false, false, false, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false,
}

func SetHelloWorld(s *Screen) {
	for i, v := range HELLO_WORLD_PIXELS {
		if !v {
			continue
		}

		x, y := indexToCoordinates(int32(i))
		s.SetPixelOn(x, y)
	}
}
