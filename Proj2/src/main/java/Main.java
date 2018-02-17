import org.lwjgl.glfw.*;
import org.lwjgl.opengl.*;
import simplex3d.algorithm.noise.*;

import static org.lwjgl.glfw.Callbacks.*;
import static org.lwjgl.glfw.GLFW.*;
import static org.lwjgl.opengl.GL11.*;
import static org.lwjgl.system.MemoryUtil.*;

public class Main {
    // The window handle
    private long window;
    private final int WIDTH = 1000;
    private final int HEIGHT = 1000;
    private final NoiseGen noise = new ClassicalGradientNoise(0);

    public void run() {
        init();
        loop();

        // Free the window callbacks and destroy the window
        glfwFreeCallbacks(window);
        glfwDestroyWindow(window);

        // Terminate GLFW and free the error callback
        glfwTerminate();
        glfwSetErrorCallback(null).free();
    }

    private void init() {
        GLFWErrorCallback.createPrint(System.err).set();

        if ( !glfwInit() )
            throw new IllegalStateException("Unable to initialize GLFW");

        glfwDefaultWindowHints(); // optional, the current window hints are already the default
        glfwWindowHint(GLFW_RESIZABLE, GLFW_FALSE); // the window will be resizable

        window = glfwCreateWindow(WIDTH, HEIGHT, "Hail Mary", NULL, NULL);
        if ( window == NULL )
            throw new RuntimeException("Failed to create the GLFW window");

        glfwMakeContextCurrent(window);

        // Enable v-sync
        glfwSwapInterval(1);
    }

    private void loop() {
        // LWJGL specific
        GL.createCapabilities();

        glClearColor(0.0f, 0.0f, 0.0f, 1.0f);

        while ( !glfwWindowShouldClose(window) ) {
            glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT); // clear the framebuffer

            glfwSwapBuffers(window);
            glfwPollEvents();
        }
    }

    private void ShaderStuff() {
        //see org.lwjgl.opengl.ARBShaderObjects
        //https://stackoverflow.com/questions/8825459/glsl-lwjgl-woes-shaders-do-nothing
    }
    public static void main(String[] args) {
        new Main().run();
    }

}