package glw

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"log"
)

const (
	DRAW_MODE_NORMAL = 0
	DRAW_MODE_WIREFRAME = 1
	DRAW_MODE_POINTS = 2
)

type GlObjDescriptor struct {
	vertexPositions, vertexColours, normals []float32
	vbo, cVbo, nbo                          uint32
}

type GlObject interface {
	GetDescriptor() *GlObjDescriptor
	Rebind()
}

func drawGlObject(obj GlObject) {
	drawGlObjectWithMode(obj, DRAW_MODE_NORMAL)
}

func drawGlObjectWithMode(obj GlObject, drawMode int32) {
	desc := obj.GetDescriptor()
	DumpErrors()

	log.Println("BIND 1")
	gl.BindBuffer(gl.ARRAY_BUFFER, desc.vbo)
	FatalDumpErrors()
	gl.EnableVertexAttribArray(0)
	FatalDumpErrors()
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	FatalDumpErrors()
	DumpErrors()

	log.Println("BIND 2")
	gl.BindBuffer(gl.ARRAY_BUFFER, desc.cVbo)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 4, gl.FLOAT, false, 0, nil)
	DumpErrors()

	log.Println("BIND 3")
	gl.BindBuffer(gl.ARRAY_BUFFER, desc.nbo)
	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)
	DumpErrors()

	gl.PointSize(3)

	if drawMode == DRAW_MODE_WIREFRAME {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	} else {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	}
	log.Println("AFTER POLYMODE")
	DumpErrors()

	points := int32(len(*desc.vertexPositions)) // For vec3s per vertex
	if drawMode == DRAW_MODE_POINTS {
		gl.DrawArrays(gl.POINTS, 0, points)
	} else {
		gl.DrawArrays(gl.TRIANGLES, 0, points)
	}
	log.Println("AFTER DRAWARRAYS")
	DumpErrors()
}
