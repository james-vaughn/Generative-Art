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

var (
	seed int64 = int64(time.Now().UTC().UnixNano())
	verts []float32 = make([]float32, 1.5*2*WIDTH*HEIGHT*3)
	noiseGen *noise.Noise = noise.NewWithSeed(seed)
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

	z := 0.0

	makePoints(z)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.DYNAMIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(program.programHandle)
	
		z += .01
		makePoints(z)
		gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(verts)*4, gl.Ptr(verts))
		gl.BindVertexArray(vao)
		gl.DrawArrays(gl.TRIANGLE_STRIP, 0, int32(len(verts) / 3))
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

//http://www.learnopengles.com/tag/triangle-strips/
func makePoints (z float64) {
	scale := float32(1.5)

	x_incr := (float32(2.0) / float32(WIDTH)) + .002
	y_incr := (float32(2.0) / float32(HEIGHT)) + .002

	idx := 0
	for y := float32(-1.0 + y_incr); y < 1.0 - y_incr; y += y_incr {
		//degenerate beginning triangle
		x_val := float64(scale * -1.0)
		y_val := float64(scale * y)		

		verts[idx] = float32(-1.0 + x_incr)
		verts[idx+1] = y
		verts[idx+2] = float32(noiseGen.Eval3(x_val, y_val, z))
		idx += 3

		for x := float32(-1.0 + x_incr); x < 1.0; x += x_incr {

			x_val := float64(scale * x)
			y_val1 := float64(scale * y)
			y_val2 := float64(scale * (y + y_incr))
			
			verts[idx] = x
			verts[idx+1] = y
			verts[idx+2] = float32(noiseGen.Eval3(x_val, y_val1, z))

			verts[idx+3] = x
			verts[idx+4] = y + y_incr
			verts[idx+5] = float32(noiseGen.Eval3(x_val, y_val2, z))
		
			idx += 6
		}

		//degenerate end triangle
		verts[idx] = verts[idx - 3]
		verts[idx + 1] = verts[idx - 2]
		verts[idx + 2] = verts[idx - 1]
		idx += 3
	}
}
