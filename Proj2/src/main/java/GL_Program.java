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
        programHandle = GL20.glCreateProgram();

        if (programHandle == 0) {
            throw new RuntimeException("Program could not be made.");
        }

        createShaderHandles(vertexShaderFile, fragmentShaderFile);
        GL20.glLinkProgram(programHandle);
        GL20.glValidateProgram(programHandle);
    }

    public void Destroy() {
        GL20.glDeleteProgram(vertexHandle);
        GL20.glDeleteProgram(fragmentHandle);
        GL20.glDeleteProgram(programHandle);
    }

    public int getProgramHandle() {
        return programHandle;
    }

    private void createShaderHandles(String vertexShaderFile, String fragmentShaderFile) throws IOException {
        vertexHandle = GL20.glCreateShader(GL20.GL_VERTEX_SHADER);
        Path path = Paths.get(vertexShaderFile);
        String vertexProg = new String(Files.readAllBytes(path));
        GL20.glShaderSource(vertexHandle, vertexProg);
        GL20.glCompileShader(vertexHandle);

        fragmentHandle = GL20.glCreateShader(GL20.GL_FRAGMENT_SHADER);
        String fragmentProg = new String(Files.readAllBytes(Paths.get(fragmentShaderFile)));
        GL20.glShaderSource(fragmentHandle, fragmentProg);
        GL20.glCompileShader(fragmentHandle);
    }
}
