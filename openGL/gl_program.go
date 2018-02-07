package main

import (
	"strings"
	"fmt"
	"io/ioutil"
	"github.com/go-gl/gl/v4.5-core/gl"
)

type Program struct {
	programHandle uint32
	vertShaderHandle uint32
	fragShaderHandle uint32
}

// Creates a new program that pulls in shaders from the given file locations
func NewProgram(vertShaderFile, fragShaderFile string) (*Program, error) {
	vertexShaderSource, _ := sourceFromFile(vertShaderFile)
	fragmentShaderSource, _ := sourceFromFile(fragShaderFile)
	
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	programHandle := gl.CreateProgram()

	gl.AttachShader(programHandle, vertexShader)
	gl.AttachShader(programHandle, fragmentShader)
	gl.LinkProgram(programHandle)

	var status int32
	gl.GetProgramiv(programHandle, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(programHandle, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(programHandle, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	program := &Program{programHandle, vertexShader, fragmentShader}
	return program, nil
}

func (p *Program) Delete() {
	fmt.Println("Deleting program")
	gl.DeleteShader(p.vertShaderHandle)
	gl.DeleteShader(p.fragShaderHandle)
	gl.DeleteProgram(p.programHandle)
}


func sourceFromFile(shaderSourceFile string) (string, error) {
	shader, err := ioutil.ReadFile(shaderSourceFile)
	if err != nil {
		return "", err
	}
	shaderStr := string(shader) + "\x00"
	return shaderStr, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
