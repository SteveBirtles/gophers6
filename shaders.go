package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	_ "image/png"
)

const terminator = "\x00"

var shaderProgram uint32

const vertexShader = ""			// <-- add this

const fragmentShader = ""		// <-- add this

func prepareShaders() {

	var err error

	shaderProgram, err = newShaderProgram(vertexShader, fragmentShader)
	if err != nil {
		panic(err)
	}
	gl.UseProgram(shaderProgram)

	projection := mgl32.Perspective(mgl32.DegToRad(45.0), float32(windowWidth)/windowHeight, 0.1, 5000.0)
	projectionUniform := gl.GetUniformLocation(shaderProgram, gl.Str("projection"+terminator))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	model := mgl32.Ident4()
	modelUniform := gl.GetUniformLocation(shaderProgram, gl.Str("model"+terminator))
	gl.UniformMatrix4fv(modelUniform, 1, false, &model[0])

	tex := gl.GetUniformLocation(shaderProgram, gl.Str("tex"+terminator))
	gl.Uniform1i(tex, 0)

	vertexIn := uint32(gl.GetAttribLocation(shaderProgram, gl.Str("vertexIn"+terminator)))
	gl.EnableVertexAttribArray(vertexIn)
	gl.VertexAttribPointer(vertexIn, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))

	texCoordIn := uint32(gl.GetAttribLocation(shaderProgram, gl.Str("texCoordIn"+terminator)))
	gl.EnableVertexAttribArray(texCoordIn)
	gl.VertexAttribPointer(texCoordIn, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))

	colourIn := uint32(gl.GetAttribLocation(shaderProgram, gl.Str("colourIn"+terminator)))
	gl.EnableVertexAttribArray(colourIn)
	gl.VertexAttribPointer(colourIn, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(5*4))

	gl.BindFragDataLocation(shaderProgram, 0, gl.Str("colourOut"+terminator))

}
