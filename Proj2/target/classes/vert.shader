#version 120 core
//version 120 for windows
//version 450 for linux

in vec3 vp;
uniform mat4 camera;
uniform mat4 ortho;

void main(void)
{
   //gl_Position = ortho * camera * vec4(vp, 1.0);
   gl_Position = vec4(vp, 1.0);
}
