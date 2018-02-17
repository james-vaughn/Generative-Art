import org.lwjgl.opengl.*;
import simplex3d.algorithm.noise.ClassicalGradientNoise;
import simplex3d.algorithm.noise.NoiseGen;

//http://wiki.lwjgl.org/wiki/Using_Vertex_Buffer_Objects_(VBO).html
public class VertexGen {
    private final NoiseGen noise = new ClassicalGradientNoise(0);
    private GL_Program program;
    private int VAOHandle;
    private int VBOHandle;
    private float[] vertices;
    private int width;
    private int height;

    private float z = 0.0f;

    public VertexGen(GL_Program program, int width, int height) {
        this.program = program;
        this.width = width;
        this.height = height;

        vertices = new float[2 * 3 * width * height];
    }

    public void GenAndBindBuffers() {
        VAOHandle = GL30.glGenVertexArrays();
        GL30.glBindVertexArray(VAOHandle);

        VBOHandle = GL15.glGenBuffers();
        GL15.glBindBuffer(GL15.GL_ARRAY_BUFFER, VBOHandle);
        GL15.glBufferData(GL15.GL_ARRAY_BUFFER, genVertices(z), GL15.GL_STATIC_DRAW);

        int inVertexLocation = GL20.glGetAttribLocation(program.getProgramHandle(), "vp");
        GL20.glEnableVertexAttribArray(inVertexLocation);
        GL20.glVertexAttribPointer(inVertexLocation, 3, GL11.GL_FLOAT, false, 0, 0);
    }

    private float[] genVertices(float z) {
        float scale = 1.5f;

        float x_incr = 4.0f * ((2.0f / (float) width));
        float y_incr = 4.0f * ((2.0f / (float) height));

        int idx = 0;
        for (float y = -1.5f + y_incr; y < 1.5f - y_incr; y += y_incr) {
            //degenerate beginning triangle
            float x_val = scale * -1.0f;
            float y_val = scale * y;

            vertices[idx] = -1.0f + x_incr;
            vertices[idx+1] = y;
            vertices[idx+2] = .9f * (float)noise.apply(x_val, y_val, z);
            idx += 3;

            for (float x = -1.5f + x_incr; x < 1.5f; x += x_incr) {

                x_val = scale * x;
                float y_val1 = scale * y;
                float y_val2 = scale * (y + y_incr);

                vertices[idx] = x;
                vertices[idx+1] = y;
                vertices[idx+2] = .9f * (float)noise.apply(x_val, y_val1, z);

                vertices[idx+3] = x;
                vertices[idx+4] = y + y_incr;
                vertices[idx+5] = .9f * (float)noise.apply(x_val, y_val2, z);

                idx += 6;
            }

            //degenerate end triangle
            vertices[idx] = vertices[idx - 3];
            vertices[idx + 1] = vertices[idx - 2];
            vertices[idx + 2] = vertices[idx - 1];
            idx += 3;
        }

        return vertices;
    }

    public void render() {
        z += .01;
        genVertices(z);

        GL15.glBufferSubData(GL15.GL_ARRAY_BUFFER, 0, vertices);
        GL30.glBindVertexArray(VAOHandle);
        GL11.glDrawArrays(GL11.GL_TRIANGLE_STRIP, 0, vertices.length/3);
    }

    public void Destroy() {
        GL15.glDeleteBuffers(VBOHandle);
    }
}
