<<<<<<< HEAD
#version 120 core
//version 120 for windows
=======
#version 450 core
//version 120 for windows
//version 450 for linux
>>>>>>> ff931ccd791f76c0045cc105939a0a42aa91a008

in vec3 vp;
uniform mat4 camera;
uniform mat4 ortho;

void main(void)
{
   //gl_Position = ortho * camera * vec4(vp, 1.0);
   gl_Position = vec4(vp, 1.0);
}
