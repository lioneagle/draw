package sdl

import (
	"syscall"
	//"unsafe"
	"errors"
	"fmt"
)

var (
	// Library
	libsdl32 = syscall.NewLazyDLL("sdl2_64.dll")
	//libmsimg32 = syscall.NewLazyDLL("msimg32.dll")

	// Functions
	procSDL_Init          = libsdl32.NewProc("SDL_Init")
	procSDL_Quit          = libsdl32.NewProc("SDL_Quit")
	procSDL_WasInit       = libsdl32.NewProc("SDL_WasInit")
	procSDL_InitSubSystem = libsdl32.NewProc("SDL_InitSubSystem")
	procSDL_QuitSubSystem = libsdl32.NewProc("SDL_QuitSubSystem")
)

func SDL_Init(flags uint32) error {
	ret, _, _ := procSDL_Init.Call(
		uintptr(flags))

	if int(ret) == -1 {
		return errors.New(fmt.Sprintf("SDL_Init failed with return %d", ret))
	}

	return nil
}

func SDL_Quit() {
	procSDL_Quit.Call()
}

func SDL_WasInit(flags uint32) error {
	ret, _, _ := procSDL_WasInit.Call(
		uintptr(flags))

	if ret == 0 {
		return errors.New(fmt.Sprintf("SDL_WasInit failed with return %d", ret))
	}

	return nil
}

func SDL_InitSubSystem(flags uint32) error {
	ret, _, _ := procSDL_InitSubSystem.Call(uintptr(flags))

	if int(ret) == -1 {
		return errors.New(fmt.Sprintf("SDL_InitSubSystem failed with return %d", ret))
	}

	return nil
}

func SDL_QuitSubSystem(flags uint32) {
	procSDL_QuitSubSystem.Call(uintptr(flags))
}
