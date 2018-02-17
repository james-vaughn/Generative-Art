<<<<<<< HEAD
#version 120 core
//version 120 for windows
=======
#version 450 core
//version 120 for windows
//version 450 for linux
>>>>>>> ff931ccd791f76c0045cc105939a0a42aa91a008

out vec4 color;
precision highp float; 

void main(void)
{
	float ratio = gl_FragCoord.z / gl_FragCoord.w;

	//colored version
	vec3 color_no_depth = vec3(0.2, 0.9* (1-ratio), 0.9*ratio);

	//depth illusion
	//vec3 color_no_depth = vec3(ratio, ratio, ratio);

	color = vec4(color_no_depth, 1.0);
}
