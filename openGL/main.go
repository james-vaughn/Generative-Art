package main

import (
	"log"
	"runtime"
	"time"
	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	noise "github.com/ojrac/opensimplex-go"
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
	gl.ClearColor(1.0, .2, .2, 1.0)

	//vertex crap
	//Pull this out into a different file

	vs := makePoints()

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vs)*4, gl.Ptr(vs), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program.programHandle)
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLE_STRIP, 0, int32(len(vs) / 3))
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

func makePoints () []float32{
	verts := make([]float32, 0)

	seed := int64(time.Now().UTC().UnixNano())
	noiseGen := noise.NewWithSeed(seed)
	
	log.Printf("Seed for noise: %d", seed)

	scale := float32(1.5)

	x_incr := float32(2.0) / float32(WIDTH)
	y_incr := float32(2.0) / float32(HEIGHT)
	for y := float32(-1.0); y < 1.0 - y_incr; y+= y_incr {
		for x := float32(-1.0); x < 1.0; x+= x_incr {
			x_val := float64(scale * x)
			y_val1 := float64(scale * y)
			y_val2 := float64(scale * (y + y_incr))
			verts = append(verts, x, y, float32(noiseGen.Eval2(x_val, y_val1)))
			verts = append(verts, x, y + y_incr, float32(noiseGen.Eval2(x_val, y_val2)))
		}
	}

	return verts
}
