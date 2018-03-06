using System;

public class ParticleColor
{
	private Random rand;
	public float R;
	public float G;
	public float B;
	public float A;

	public ParticleColor ()
	{
		rand = new Random(Guid.NewGuid().GetHashCode());
	}

	public ParticleColor SetRandomColor(float alpha) {
		R = (float)rand.NextDouble ();
		G = (float)rand.NextDouble ();
		B = (float)rand.NextDouble ();
		A = alpha;

		return this;
	}

	public ParticleColor SetRandomGrayColor(float alpha) {
		R = (float)rand.NextDouble ();
		G = R;
		B = R;
		A = alpha;

		return this;
	}



	public ParticleColor Mutate() {
		var color = (float)rand.NextDouble ();
		var sign = (float)rand.NextDouble () > .5f ? -1f : 1f;
		var amount = (float)rand.NextDouble() * .03f;

		if (color < .33f) {
			R += sign * amount;
			R = clamp (R);
		} else if (color < .67f) {
			G += sign * amount;
			G = clamp (G);
		} else {
			B += sign * amount;
			B = clamp (B);
		}

		return this;
	}

	public ParticleColor MutateGray() {
		var sign = (float)rand.NextDouble () > .5f ? -1f : 1f;
		var amount = (float)rand.NextDouble() * .03f;

		R += sign * amount;
		R = clamp (R);
		G += sign * amount;
		G = clamp (G);
		B += sign * amount;
		B = clamp (B);

		return this;
	}

	private float clamp(float val) {
		return val > 1f ? 1f : val < 0f ? 0f : val;
	}
}

