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

var verts = []float32 {
	//1st triangle
	-.4, 0.1, -1,
	0.0, 1.0, 0,
	.4, 0.1, 1,

	//2nd triangle
	-.4, -0.1, -1,
	0.0, -1.0, 0,
	.4, -0.1, 1,

}

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

	program, err := NewProgram("vert.shader", "frag.shader")
	if err != nil {
		log.Fatal(err)
	}
	
	defer program.Delete()

	gl.UseProgram(program.programHandle)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(.2, .2, .2, 1.0)

	//vertex crap
	//Pull this out into a different file
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.STATIC_DRAW)

	//vertAttrib := uint32(gl.GetAttribLocation(program.programHandle, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program.programHandle)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
		window.SwapBuffers()
		glfw.PollEvents()
	}

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
