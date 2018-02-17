#version 120 core
//version 120 for windows
//version 450 for linux

out vec4 color;
precision highp float; 

void main(void)
{
	float ratio = gl_FragCoord.z / gl_FragCoord.w;

	//colored version
	//vec3 color_no_depth = vec3(0.2, 0.9* (1-ratio), 0.9*ratio);

	//depth illusion
	vec3 color_no_depth = vec3(ratio, ratio, ratio);

	color = vec4(color_no_depth, 1.0);
}
