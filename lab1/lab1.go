package main

// This initial env is based on
// https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl
// and Iain Martin's basic_wrapper example C++ code.

import (
	"runtime"
	"log"
	"github.com/emberwalker/ac41001_opengl/glw"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	height = 750
	width  = 750
)

var (
	matrix = []float32 {
		.75, .75, 0, 1,
		.75, -.75, 0, 1,
		-.75, -.75, 0, 1,
	}

	colorMatrix = []float32 {
		1.0, 0.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
	}
)

func main() {
	runtime.LockOSThread()
	log.Println("Lab 1: init")

	glfwWrapper := glw.NewGlfwWrapper(width, height, 60.0, "Lab 1")
	program := glw.InitOpenGl()

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4 * len(matrix), gl.Ptr(matrix), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	var cVbo uint32
	gl.GenBuffers(1, &cVbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, cVbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4 * len(colorMatrix), gl.Ptr(colorMatrix), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 0, nil)

	fragShader, err := glw.CompileShader("lab1/frag.glsl", gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	vertShader, err := glw.CompileShader("lab1/vert.glsl", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	gl.AttachShader(program, fragShader)
	gl.AttachShader(program, vertShader)
	gl.LinkProgram(program)

	var x float32 = 0.0
	var y float32 = .75
	var hInc float32 = 0.00
	var vInc float32 = 0.00

	glfwWrapper.KeyCallback = func(wrapper *glw.GlfwWrapper, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		hInc = 0.0
		vInc = 0.0
		if key == glfw.KeyEscape && action == glfw.Press {
			wrapper.Window.SetShouldClose(true)
		}

		// Horizontal motion
		if key == glfw.KeyLeft && (action == glfw.Press || action == glfw.Repeat) {
			hInc = -0.01
		} else if key == glfw.KeyRight && (action == glfw.Press || action == glfw.Repeat) {
			hInc = 0.01
		}

		// Vertical motion
		if key == glfw.KeyDown && (action == glfw.Press || action == glfw.Repeat) {
			vInc = -0.01
		} else if key == glfw.KeyUp && (action == glfw.Press || action == glfw.Repeat) {
			vInc = 0.01
		}
	}

	glfwWrapper.Renderer = func(wrapper *glw.GlfwWrapper) {
		matrix[0] = x
		matrix[1] = y

		gl.BufferData(gl.ARRAY_BUFFER, 4 * len(matrix), gl.Ptr(matrix), gl.DYNAMIC_DRAW)

		gl.ClearColor(1, 1, 1, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.UseProgram(program)

		gl.BindBuffer(gl.ARRAY_BUFFER, cVbo)
		gl.EnableVertexAttribArray(1)
		gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 0, nil)

		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.EnableVertexAttribArray(0)
		gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 0, nil)

		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		gl.DisableVertexAttribArray(0)
		gl.DisableVertexAttribArray(1)
		gl.UseProgram(0)

		x += hInc
		y += vInc
	}

	glfwWrapper.EventLoop()
}
