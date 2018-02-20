using System.Collections;
using System.Collections.Generic;
using UnityEngine;

//https://docs.unity3d.com/ScriptReference/ParticleSystem.html
//https://docs.unity3d.com/ScriptReference/ParticleSystem.TrailModule.html
public class FlowField : MonoBehaviour
{
	private ParticleSystem ps;
	private Gradient gradient = new Gradient();
	public bool swapColors = true;

	void Start()
	{
		gradient.SetKeys(
			new GradientColorKey[] { new GradientColorKey(Color.blue, 0.0f), new GradientColorKey(Color.green, 1.0f) },
			new GradientAlphaKey[] { new GradientAlphaKey(1.0f, 0.0f), new GradientAlphaKey(1.0f, 1.0f) }
		);

		ps = GetComponent<ParticleSystem>();

		var trails = ps.trails;
		trails.enabled = true;

		var psr = GetComponent<ParticleSystemRenderer>();
		psr.trailMaterial = new Material(Shader.Find("Sprites/Default"));
		trails.colorOverTrail = gradient;
	}
}