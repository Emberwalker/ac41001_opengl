#version 410

out vec4 outputColor;
in vec4 colour_for_frag;
void main() {
	outputColor = colour_for_frag;
}
