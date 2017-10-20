package glw

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl32/matstack"
)

// Reverse-order queueing so we can do GL matrices in the right order.
type TransformQueue struct {
	transforms []mgl32.Mat4
}

func NewTransformQueue() *TransformQueue {
	return new(TransformQueue)
}

func (queue *TransformQueue) Push(mat4 mgl32.Mat4) {
	queue.transforms = append(queue.transforms, mat4)
}

func (queue *TransformQueue) Result() mgl32.Mat4 {
	out := mgl32.Ident4()
	elements := len(queue.transforms)
	// Map over the transform matrices in reverse order (GL-style)
	for i := range queue.transforms {
		out = out.Mul4(queue.transforms[elements - (1 + i)])
	}
	return out
}

// Extensions to matstack
type GlMatStack struct {
	*matstack.MatStack
}

func NewGlMatStack() *GlMatStack {
	return &GlMatStack{matstack.NewMatStack()}
}

func (ms *GlMatStack) Pop() {
	if err := ms.MatStack.Pop(); err != nil {
		panic(err)
	}
}

func (ms *GlMatStack) Mul(mat mgl32.Mat4) {
	ms.RightMul(mat)
}
