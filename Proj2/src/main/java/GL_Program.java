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

        vertexHandle = createShaderHandle(vertexShaderFile, GL20.GL_VERTEX_SHADER);
        fragmentHandle = createShaderHandle(fragmentShaderFile, GL20.GL_FRAGMENT_SHADER);

        if (vertexHandle == 0 || fragmentHandle == 0) {
            System.err.println("Issue creating shader handles.");
        }

        GL20.glAttachShader(programHandle, vertexHandle);
        GL20.glAttachShader(programHandle, fragmentHandle);
        LinkProgram();
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

    private int createShaderHandle(String shaderFile, int shaderType) throws IOException {
        int handle = GL20.glCreateShader(shaderType);
        Path path = Paths.get(shaderFile);
        String prog = new String(Files.readAllBytes(path));
        GL20.glShaderSource(handle, prog);
        GL20.glCompileShader(handle);

        int[] status = new int[1];
        GL20.glGetShaderiv(handle, GL20.GL_COMPILE_STATUS, status);
        if (status[0] == GL11.GL_FALSE) {
            System.err.println("Issue compiling shaders:");
            System.err.println(GL20.glGetShaderInfoLog(handle));
        }

        return handle;
    }

    private void LinkProgram() {
        GL20.glLinkProgram(programHandle);

        int[] status = new int[1];
        GL20.glGetProgramiv(programHandle, GL20.GL_LINK_STATUS, status);
        if (status[0] == GL11.GL_FALSE) {
            System.err.println("Issue linking shaders to program:");
            System.err.println(GL20.glGetProgramInfoLog(programHandle));
        }
    }
}
