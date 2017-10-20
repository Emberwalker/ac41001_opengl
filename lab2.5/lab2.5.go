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
	//"github.com/go-gl/mathgl/mgl32"
)

const (
	height = 750
	width  = 750

	increment = 0.05
)

var (
	matrix = []float32 {
		-0.25, 0.25, -0.25, 1,
		-0.25, -0.25, -0.25, 1,
		0.25, -0.25, -0.25, 1,

		0.25, -0.25, -0.25, 1,
		0.25, 0.25, -0.25, 1,
		-0.25, 0.25, -0.25, 1,

		0.25, -0.25, -0.25, 1,
		0.25, -0.25, 0.25, 1,
		0.25, 0.25, -0.25, 1,

		0.25, -0.25, 0.25, 1,
		0.25, 0.25, 0.25, 1,
		0.25, 0.25, -0.25, 1,

		0.25, -0.25, 0.25, 1,
		-0.25, -0.25, 0.25, 1,
		0.25, 0.25, 0.25, 1,

		-0.25, -0.25, 0.25, 1,
		-0.25, 0.25, 0.25, 1,
		0.25, 0.25, 0.25, 1,

		-0.25, -0.25, 0.25, 1,
		-0.25, -0.25, -0.25, 1,
		-0.25, 0.25, 0.25, 1,

		-0.25, -0.25, -0.25, 1,
		-0.25, 0.25, -0.25, 1,
		-0.25, 0.25, 0.25, 1,

		-0.25, -0.25, 0.25, 1,
		0.25, -0.25, 0.25, 1,
		0.25, -0.25, -0.25, 1,

		0.25, -0.25, -0.25, 1,
		-0.25, -0.25, -0.25, 1,
		-0.25, -0.25, 0.25, 1,

		-0.25, 0.25, -0.25, 1,
		0.25, 0.25, -0.25, 1,
		0.25, 0.25, 0.25, 1,

		0.25, 0.25, 0.25, 1,
		-0.25, 0.25, 0.25, 1,
		-0.25, 0.25, -0.25, 1,
	}

	colorMatrix = []float32 {
		0.0, 0.0, 1.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
		0.0, 0.0, 1.0, 1.0,
		0.0, 0.0, 1.0, 1.0,

		0.0, 1.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
		0.0, 1.0, 0.0, 1.0,

		1.0, 1.0, 0.0, 1.0,
		1.0, 1.0, 0.0, 1.0,
		1.0, 1.0, 0.0, 1.0,
		1.0, 1.0, 0.0, 1.0,
		1.0, 1.0, 0.0, 1.0,
		1.0, 1.0, 0.0, 1.0,

		1.0, 0.0, 0.0, 1.0,
		1.0, 0.0, 0.0, 1.0,
		1.0, 0.0, 0.0, 1.0,
		1.0, 0.0, 0.0, 1.0,
		1.0, 0.0, 0.0, 1.0,
		1.0, 0.0, 0.0, 1.0,

		1.0, 0.0, 1.0, 1.0,
		1.0, 0.0, 1.0, 1.0,
		1.0, 0.0, 1.0, 1.0,
		1.0, 0.0, 1.0, 1.0,
		1.0, 0.0, 1.0, 1.0,
		1.0, 0.0, 1.0, 1.0,

		0.0, 1.0, 1.0, 1.0,
		0.0, 1.0, 1.0, 1.0,
		0.0, 1.0, 1.0, 1.0,
		0.0, 1.0, 1.0, 1.0,
		0.0, 1.0, 1.0, 1.0,
		0.0, 1.0, 1.0, 1.0,
	}
)

func main() {
	runtime.LockOSThread()
	log.Println("Lab 2: init")

	glfwWrapper := glw.NewGlfwWrapper(width, height, 60.0, "Lab 2")
	program := glw.InitOpenGl()
	glw.FatalDumpErrors()

	/*var vbo uint32
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
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 0, nil)*/

	fragShader, err := glw.CompileShader("lab2.5/frag.glsl", gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}
	vertShader, err := glw.CompileShader("lab2.5/vert.glsl", gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	glw.FatalDumpErrors()

	gl.AttachShader(program, fragShader)
	gl.AttachShader(program, vertShader)
	gl.LinkProgram(program)
	gl.Enable(gl.DEPTH_TEST)
	glw.FatalDumpErrors()

	modelId := gl.GetUniformLocation(program, gl.Str("model" + "\x00"))
	//altCalcUniform := gl.GetUniformLocation(program, gl.Str("altCalc" + "\x00"))

	var angleX float32 = 0.0
	var angleXInc float32 = 0.0

	var angleY float32 = 0.0
	var angleYInc float32 = 0.0

	var x, y, z float32 = 0, 0, 0
	var scaleX, scaleY, scaleZ float32 = 1, 1, 1

	useAltColourCalc := false

	glfwWrapper.KeyCallback = func(wrapper *glw.GlfwWrapper, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action != glfw.Press { return }
		if key == glfw.KeyEscape {
			wrapper.Window.SetShouldClose(true)
		}

		// Rotation keys
		if key == glfw.KeyQ {
			angleXInc += increment
		}
		if key == glfw.KeyE {
			angleXInc -= increment
		}
		if key == glfw.KeyZ {
			angleYInc += increment
		}
		if key == glfw.KeyX {
			angleYInc -= increment
		}

		// Scale keys
		if key == glfw.KeyL {
			scaleX += increment
		}
		if key == glfw.KeyH {
			scaleX -= increment
		}
		if key == glfw.KeyK {
			scaleY += increment
		}
		if key == glfw.KeyJ {
			scaleY -= increment
		}
		if key == glfw.KeyI {
			scaleZ += increment
		}
		if key == glfw.KeyM {
			scaleZ -= increment
		}

		// Movement keys
		if key == glfw.KeyW {
			y += increment
		}
		if key == glfw.KeyS {
			y -= increment
		}
		if key == glfw.KeyD {
			x += increment
		}
		if key == glfw.KeyA {
			x -= increment
		}
		if key == glfw.KeyR {
			z += increment
		}
		if key == glfw.KeyF {
			z -= increment
		}

		// Toggles
		if key == glfw.KeySpace {
			useAltColourCalc = !useAltColourCalc
		}
	}

	cube := glw.NewGlCube()
	cube2 := glw.NewGlCube()
	log.Println(cube2)

	glfwWrapper.Renderer = func(wrapper *glw.GlfwWrapper) {
		gl.ClearColor(1, 1, 1, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		gl.UseProgram(program)

		transformQueue := glw.NewTransformQueue()
		/*transformQueue.Push(mgl32.Scale3D(scaleX, scaleY, scaleZ))
		transformQueue.Push(mgl32.HomogRotate3DX(-angleX))
		transformQueue.Push(mgl32.HomogRotate3DY(-angleY))
		transformQueue.Push(mgl32.Translate3D(x, y, z))*/

		model := transformQueue.Result()
		gl.UniformMatrix4fv(modelId, 1, false, &model[0])

		/*var shouldUseAltCalcInt int32 = 0
		if useAltColourCalc { shouldUseAltCalcInt = 1 }
		gl.Uniform1i(altCalcUniform, shouldUseAltCalcInt)*/

		//gl.DrawArrays(gl.TRIANGLES, 0, 36)
		cube.DrawWithMode(glw.DRAW_MODE_POINTS)

		gl.UseProgram(0)

		angleX += angleXInc
		angleY += angleYInc
	}

	glfwWrapper.EventLoop()
}
