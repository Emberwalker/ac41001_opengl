// Vertex shader with a model transformation uniform

#version 400 core

layout(location = 0) in vec3 rawPosition;
layout(location = 1) in vec4 colour;
layout(location = 2) in vec3 normal;
out vec4 fcolour;
uniform mat4 model;
//uniform int altCalc;

void main()
{
    vec4 position = vec4(rawPosition, 1.0);
	//gl_Position = model * position;
	gl_Position = position;

	//if (altCalc == 0) {
	    fcolour = colour;
	//} else {
	//    fcolour = position * 2.0 + vec4(0.5, 0.5, 0.5, 1.0);
	//}
}