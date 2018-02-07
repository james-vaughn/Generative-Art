package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	WIDTH  int = 1000
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
		log.Fatal(err)
	}

	defer window.Destroy()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		log.Fatal(err)
	}

	//program, err := NewProgram("vert.shader", "frag.shader")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer program.Delete()

}

func createWindow() (*glfw.Window, error) {
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 5)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(WIDTH, HEIGHT, "Hail Mary", nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()

	return window, nil

}
