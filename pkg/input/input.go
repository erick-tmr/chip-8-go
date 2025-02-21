package input

import (
	"fmt"

	"github.com/erick-tmr/chip-8-go/pkg/emulator"
	"github.com/veandco/go-sdl2/sdl"
)

// Commands
type Command interface {
	Execute(emulatorInstance *emulator.Emulator)
}

type QuitCommand struct{}

func (c QuitCommand) Execute(emulatorInstance *emulator.Emulator) {
	emulatorInstance.State = emulator.Quitted
}

type TogglePauseCommand struct{}

func (c TogglePauseCommand) Execute(emulatorInstance *emulator.Emulator) {
	if emulatorInstance.State == emulator.Running {
		emulatorInstance.State = emulator.Paused
		fmt.Println("emulator paused")
	} else {
		emulatorInstance.State = emulator.Running
		fmt.Println("emulator resumed")
	}
}

// Event Handlers
type EventHandler interface {
	HandleEvent(event sdl.Event) (Command, bool)
}

type QuitEventHandler struct {
	command Command
}

func NewQuitEventHandler() *QuitEventHandler {
	return &QuitEventHandler{command: QuitCommand{}}
}

func (h *QuitEventHandler) HandleEvent(event sdl.Event) (Command, bool) {
	return h.command, true
}

type KeyboardEventHandler struct {
	keyCommands map[sdl.Keycode]Command
}

func NewKeyboardEventHandler() *KeyboardEventHandler {
	return &KeyboardEventHandler{
		keyCommands: map[sdl.Keycode]Command{
			sdl.K_SPACE: TogglePauseCommand{},
		},
	}
}

func (h *KeyboardEventHandler) HandleEvent(event sdl.Event) (Command, bool) {
	coarsedEvent := event.(*sdl.KeyboardEvent)
	cmd, ok := h.keyCommands[coarsedEvent.Keysym.Sym]

	return cmd, ok
}

// Manages all inputs
type InputHandler struct {
	handlers map[uint32]EventHandler
}

func NewInputHandler() *InputHandler {
	handler := &InputHandler{
		handlers: map[uint32]EventHandler{},
	}

	// Register default handlers
	handler.handlers[sdl.QUIT] = NewQuitEventHandler()
	handler.handlers[sdl.KEYDOWN] = NewKeyboardEventHandler()

	return handler
}

func (i *InputHandler) HandleInput(emulatorInstance *emulator.Emulator) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		handler, exists := i.handlers[event.GetType()]
		if !exists {
			continue
		}

		cmd, exists := handler.HandleEvent(event)
		if !exists {
			continue
		}

		cmd.Execute(emulatorInstance)
		return
	}
}
