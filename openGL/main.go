package main

import (
	"runtime"
	"log"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	WIDTH int = 1000
	HEIGHT int = 1000
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	window, err := createWindow()
	if err != nil {
		panic(err)
	}

	defer window.Destroy()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	program, err := NewProgram("", "")
	if err != nil {
		panic(err)
	}

	defer program.Delete()

}


func createWindow() (*glfw.Window, error) {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(WIDTH, HEIGHT, "Hail Mary", nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()

	return window, nil

}
