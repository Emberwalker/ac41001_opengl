#version 410

layout(location = 0) in vec4 position;
layout(location = 1) in vec4 colour;
out vec4 colour_for_frag;
void main() {
	gl_Position = position;
	colour_for_frag = colour;
}
