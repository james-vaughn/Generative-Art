import org.lwjgl.opengl.*;

//http://wiki.lwjgl.org/wiki/Using_Vertex_Buffer_Objects_(VBO).html
public class VertexGen {
    private GL_Program program;
    private int VAOHandle;
    private int VBOHandle;

    public VertexGen(GL_Program program) {
        this.program = program;

        VAOHandle = ARBVertexArrayObject.glGenVertexArrays();
        GL30.glBindVertexArray(VAOHandle);

        VBOHandle = GL15.glGenBuffers();
        GL15.glBindBuffer(GL15.GL_ARRAY_BUFFER, VBOHandle);
        GL15.glBufferData(VBOHandle, genVertices(), GL15.GL_DYNAMIC_DRAW);

        GL20.glEnableVertexAttribArray(0);
        GL20.glVertexAttribPointer(0, 3, GL11.GL_FLOAT, false, 0, 0);
    }

    private float[] genVertices() {
        return new float[] {
                -1.0f, -1.0f, 0.0f,
                0.0f, 1.0f, 0.0f,
                1.0f, -1.0f, 0.0f
        };
    }

    public void bindVertexArrays() {
        GL30.glBindVertexArray(VAOHandle);
    }
}
