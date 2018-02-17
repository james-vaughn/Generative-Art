import org.lwjgl.opengl.*;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

public class GL_Program {
    private int programHandle;
    private int vertexHandle;
    private int fragmentHandle;

    public GL_Program(String vertexShaderFile, String fragmentShaderFile) throws IOException {
        programHandle = ARBShaderObjects.glCreateProgramObjectARB();

        if (programHandle == 0) {
            throw new RuntimeException("Program could not be made.");
        }

        createShaderHandles(vertexShaderFile, fragmentShaderFile);
        ARBShaderObjects.glLinkProgramARB(programHandle);
        ARBShaderObjects.glValidateProgramARB(programHandle);
    }

    public void Destroy() {
        ARBShaderObjects.glDeleteObjectARB(vertexHandle);
        ARBShaderObjects.glDeleteObjectARB(fragmentHandle);
        ARBShaderObjects.glDeleteObjectARB(programHandle);
    }

    public int getProgramHandle() {
        return programHandle;
    }

    private void createShaderHandles(String vertexShaderFile, String fragmentShaderFile) throws IOException {
        vertexHandle = ARBShaderObjects.glCreateShaderObjectARB(ARBVertexShader.GL_VERTEX_SHADER_ARB);
        Path path = Paths.get(vertexShaderFile);
        String vertexProg = new String(Files.readAllBytes(path));
        ARBShaderObjects.glShaderSourceARB(vertexHandle, vertexProg);
        ARBShaderObjects.glCompileShaderARB(vertexHandle);

        fragmentHandle = ARBShaderObjects.glCreateShaderObjectARB(ARBFragmentShader.GL_FRAGMENT_SHADER_ARB);
        String fragmentProg = new String(Files.readAllBytes(Paths.get(fragmentShaderFile)));
        ARBShaderObjects.glShaderSourceARB(fragmentHandle, fragmentProg);
        ARBShaderObjects.glCompileShaderARB(fragmentHandle);
    }
}
