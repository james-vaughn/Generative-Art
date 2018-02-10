#version 450 core
out vec4 color;
precision highp float; 

void main(void)
{
	float ratio = gl_FragCoord.z / gl_FragCoord.w;
	vec3 color_no_depth = vec3(0.2, 0.9* (1-ratio), 0.9*ratio);  
	color = vec4(color_no_depth, 1.0);
}
