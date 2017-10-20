package glw

import "github.com/go-gl/gl/v4.1-core/gl"

var (
	tmplVertexPositions = []float32{
		-0.25, 0.25, -0.25,
		-0.25, -0.25, -0.25,
		0.25, -0.25, -0.25,

		0.25, -0.25, -0.25,
		0.25, 0.25, -0.25,
		-0.25, 0.25, -0.25,

		0.25, -0.25, -0.25,
		0.25, -0.25, 0.25,
		0.25, 0.25, -0.25,

		0.25, -0.25, 0.25,
		0.25, 0.25, 0.25,
		0.25, 0.25, -0.25,

		0.25, -0.25, 0.25,
		-0.25, -0.25, 0.25,
		0.25, 0.25, 0.25,

		-0.25, -0.25, 0.25,
		-0.25, 0.25, 0.25,
		0.25, 0.25, 0.25,

		-0.25, -0.25, 0.25,
		-0.25, -0.25, -0.25,
		-0.25, 0.25, 0.25,

		-0.25, -0.25, -0.25,
		-0.25, 0.25, -0.25,
		-0.25, 0.25, 0.25,

		-0.25, -0.25, 0.25,
		0.25, -0.25, 0.25,
		0.25, -0.25, -0.25,

		0.25, -0.25, -0.25,
		-0.25, -0.25, -0.25,
		-0.25, -0.25, 0.25,

		-0.25, 0.25, -0.25,
		0.25, 0.25, -0.25,
		0.25, 0.25, 0.25,

		0.25, 0.25, 0.25,
		-0.25, 0.25, 0.25,
		-0.25, 0.25, -0.25,
	}

	tmplVertexColours = []float32{
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

	tmplNormals = []float32{
		0, 0, -1., 0, 0, -1., 0, 0, -1.,
		0, 0, -1., 0, 0, -1., 0, 0, -1.,
		1., 0, 0, 1., 0, 0, 1., 0, 0,
		1., 0, 0, 1., 0, 0, 1., 0, 0,
		0, 0, 1., 0, 0, 1., 0, 0, 1.,
		0, 0, 1., 0, 0, 1., 0, 0, 1.,
		-1., 0, 0, -1., 0, 0, -1., 0, 0,
		-1., 0, 0, -1., 0, 0, -1., 0, 0,
		0, -1., 0, 0, -1., 0, 0, -1., 0,
		0, -1., 0, 0, -1., 0, 0, -1., 0,
		0, 1., 0, 0, 1., 0, 0, 1., 0,
		0, 1., 0, 0, 1., 0, 0, 1., 0,
	}

	// Enforce at compile-time GlCube implements GlObject
	_ GlObject = (*GlCube)(nil)
)

type GlCube struct {
	descriptor *GlObjDescriptor
}

func NewGlCube() *GlCube {
	// Clear any junk errors prior
	DumpErrors()

	positions := make([]float32, len(tmplVertexPositions))
	copy(positions, tmplVertexPositions)
	colours := make([]float32, len(tmplVertexColours))
	copy(colours, tmplVertexColours)
	normals := make([]float32, len(tmplNormals))
	copy(normals, tmplNormals)

	var vbo, cVbo, nbo uint32
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &cVbo)
	gl.GenBuffers(1, &nbo)
	FatalDumpErrors()

	cube := &GlCube{
		&GlObjDescriptor{
			vertexPositions: positions,
			vertexColours:   colours,
			normals:         normals,
			vbo:             vbo,
			cVbo:            cVbo,
			nbo:             nbo,
		},
	}

	cube.Rebind()

	return cube
}

func (cube *GlCube) GetDescriptor() *GlObjDescriptor {
	return cube.descriptor
}

func (cube *GlCube) Rebind() {
	desc := cube.descriptor

	gl.BindBuffer(gl.ARRAY_BUFFER, desc.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(desc.vertexPositions), gl.Ptr(desc.vertexPositions), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, desc.cVbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(desc.vertexColours), gl.Ptr(desc.vertexColours), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, desc.nbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(desc.normals), gl.Ptr(desc.normals), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	FatalDumpErrors()
}

func (cube *GlCube) Draw() {
	drawGlObject(cube)
}

func (cube *GlCube) DrawWithMode(drawMode int32) {
	drawGlObjectWithMode(cube, drawMode)
}
