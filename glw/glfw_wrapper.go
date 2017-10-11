package glw

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

/*
 * This file contains a Go port of Iain Martin's basic_wrapper wrapper_glfw.h/cpp
 */

type GlfwWrapper struct {
	Window              *glfw.Window
	width, height       int
	Fps                 float32 // TODO: What's this for? It isn't used in basic_wrapper.
	title               string
	running, ShouldWait bool
	Renderer            func(wrapper *GlfwWrapper)
	Reshape             func(wrapper *GlfwWrapper, w, h int)
	KeyCallback         func(wrapper *GlfwWrapper, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey)
}

// NewGlfwWrapper creates a new GlfwWrapper and returns it initialised. Also calls runtime.LockOSThread() to ensure
// the context is properly setup.
func NewGlfwWrapper(width, height int, fps float32, title string) *GlfwWrapper {
	runtime.LockOSThread()
	wrapper := &GlfwWrapper{
		Window:      initGlfw(width, height, title),
		width:       width,
		height:      height,
		Fps:         fps,
		title:       title,
		running:     false,
		ShouldWait:  false,
		Renderer: 	 func(wrapper *GlfwWrapper) {},
		Reshape:     func(wrapper *GlfwWrapper, w, h int) {},
		KeyCallback: func(wrapper *GlfwWrapper, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {},
	}

	wrapper.Window.SetSizeCallback(func(window *glfw.Window, w int, h int) {
		wrapper.Reshape(wrapper, w, h)
	})

	wrapper.Window.SetKeyCallback(func(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		wrapper.KeyCallback(wrapper, key, scancode, action, mods)
	})

	return wrapper
}

// From https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl
func initGlfw(width, height int, title string) *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	//glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetInputMode(glfw.StickyKeysMode, glfw.True)

	return window
}

// Destroy destroys the GLFW context and invalidates the wrapper reference to the GLFW Window.
func (wrapper *GlfwWrapper) Destroy() {
	wrapper.Window = nil
	glfw.Terminate()
}

// SetTitle sets the window title
func (wrapper *GlfwWrapper) SetTitle(title string) {
	wrapper.title = title
	wrapper.Window.SetTitle(title)
}

func (wrapper *GlfwWrapper) EventLoop() {
	for !wrapper.Window.ShouldClose() {
		// Call user-provided renderer
		wrapper.Renderer(wrapper)

		// Swap buffers (GLFW is double-buffered, so this need swapped every frame)
		wrapper.Window.SwapBuffers()

		// Check for inputs
		if wrapper.ShouldWait {
			glfw.WaitEvents()
		} else {
			glfw.PollEvents()
		}
	}

	wrapper.Destroy()
}
