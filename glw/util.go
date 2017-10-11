package glw

// Based mostly on
// https://kylewbanks.com/blog/tutorial-opengl-with-golang-part-1-hello-opengl

import (
	"log"
	//"runtime"
	// We need to use OpenGL 4.1 as this is the highest version on macOS (whyyy)
	"github.com/go-gl/gl/v4.1-core/gl"
	"os"
	"strings"
	"io/ioutil"
	"fmt"
)

var (
	goProjectRoot = strings.Join([]string{
		os.Getenv("GOPATH"),
		"src",
		"github.com",
		"emberwalker",
		"ac41001_opengl",
		}, string(os.PathSeparator))
)

// InitOpenGl initialises OpenGL and returns an initialised program.
func InitOpenGl() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	printVersion()

	prog := gl.CreateProgram()
	return prog
}

func printVersion() {
	var major int32
	var minor int32
	gl.GetIntegerv(gl.MAJOR_VERSION, &major)
	gl.GetIntegerv(gl.MINOR_VERSION, &minor)
	vendor := gl.GoStr(gl.GetString(gl.VENDOR))
	shaderLevel := gl.GoStr(gl.GetString(gl.SHADING_LANGUAGE_VERSION))
	exts := gl.GoStr(gl.GetString(gl.EXTENSIONS))

	log.Printf("OpenGL version %d.%d (%s) with GLSL %s", major, minor, vendor, shaderLevel)
	if exts != "" {
		log.Println("Available OpenGL extensions:", strings.Replace(exts, " ", ", ", -1))
	} else {
		log.Println("No OpenGL extensions in environment.")
	}

	/*version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)*/
}

// CompileShader gets the shader in the given file (relative to the project root), compiles it with OpenGL and returns
// the shaders OpenGL name (index)
func CompileShader(fname string, shaderType uint32) (uint32, error) {
	source, err := getShaderSrcFile(fname)
	if err != nil {
		panic(err)
	}

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

		logInfo := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logInfo))

		return 0, fmt.Errorf("failed to compile %v: %v", source, logInfo)
	}

	return shader, nil
}

func getShaderSrcFile(fname string) (string, error) {
	_, exists := os.LookupEnv("GOPATH")
	if !exists {
		panic("No GOPATH defined.")
	}

	fContents, err := ioutil.ReadFile(goProjectRoot + string(os.PathSeparator) + fname)
	if err != nil {
		return "", err
	} else {
		// Don't forget your null-termination every morning, Dr OpenGL says so.
		return string(fContents) + "\x00", nil
	}
}
