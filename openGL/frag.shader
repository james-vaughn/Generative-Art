#version 450 core
out vec4 color;
 
void main(void)
{
	vec3 color_no_depth = vec3(0.0, 1.0, 1.0);  
	color = vec4(gl_FragCoord.z * color_no_depth, 1.0);
}
