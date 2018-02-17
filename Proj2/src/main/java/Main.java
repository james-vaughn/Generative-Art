import org.lwjgl.glfw.*;
import org.lwjgl.opengl.*;
import java.io.IOException;

import static org.lwjgl.glfw.Callbacks.*;
import static org.lwjgl.glfw.GLFW.*;
import static org.lwjgl.opengl.GL11.*;
import static org.lwjgl.system.MemoryUtil.*;

public class Main {
    // The window handle
    private long window;
    private final int WIDTH = 1000;
    private final int HEIGHT = 1000;
    private final boolean WINDOWS = true;

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

        //Following disabled to work on windows
        if (WINDOWS == false) {
            glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE);
            glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 4);
            glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 5);
            glfwWindowHint(GLFW_OPENGL_FORWARD_COMPAT, GLFW_TRUE);
        }

        window = glfwCreateWindow(WIDTH, HEIGHT, "Hail Mary", NULL, NULL);
        if ( window == NULL )
            throw new RuntimeException("Failed to create the GLFW window");

        glfwMakeContextCurrent(window);

        // Enable v-sync
        glfwSwapInterval(1);

        // LWJGL specific
        GL.createCapabilities();
    }

    private void loop() {
        GL_Program program;

        try {
            program = new GL_Program("src/main/resources/vert.shader",
                    "src/main/resources/frag.shader");
        } catch (IOException e) {
            throw new RuntimeException("Could not create program: "+ e.getMessage());
        }

        VertexGen vertGen = new VertexGen(program, WIDTH, HEIGHT);
        vertGen.GenAndBindBuffers();

        GL11.glEnable(GL11.GL_DEPTH_TEST);
        GL11.glDepthFunc(GL11.GL_LESS);
        glClearColor(0.0f, 0.0f, 0.0f, 0.0f);

        while ( !glfwWindowShouldClose(window) ) {
            glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);
            GL20.glUseProgram(program.getProgramHandle());

            vertGen.render();

            glfwSwapBuffers(window);
            glfwPollEvents();
        }

        vertGen.Destroy();
        program.Destroy();
    }

    public static void main(String[] args) {
        new Main().run();
    }

}