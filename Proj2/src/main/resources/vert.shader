#version 450 core
 
in vec3 vp;
uniform mat4 camera;
uniform mat4 ortho;

void main(void)
{
   gl_Position = ortho * camera * vec4(vp, 1.0);
}
