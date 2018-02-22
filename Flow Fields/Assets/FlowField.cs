using System.Collections;
using System.Collections.Generic;
using UnityEngine;

//https://docs.unity3d.com/ScriptReference/ParticleSystem.html
//https://docs.unity3d.com/ScriptReference/ParticleSystem.TrailModule.html
public class FlowField : MonoBehaviour
{
	private ParticleSystem ps;
	private ParticleSystem.Particle[] particles;
	private ParticleColor color = new ParticleColor();
	public bool swapColors = true;

	void Start()
	{
		color = color.SetRandomColor ();

		ps = GetComponent<ParticleSystem>();
		particles = new ParticleSystem.Particle[ps.main.maxParticles];
		int partCount = ps.GetParticles(particles);

		setParticleColors (partCount);
		setParticleSizes (partCount);

		var trails = ps.trails;
		trails.enabled = true;

		var psr = GetComponent<ParticleSystemRenderer>();
		psr.trailMaterial = new Material(Shader.Find("Sprites/Default"));
	}


	void Update() {
		int partCount = ps.GetParticles(particles);
		setParticleColors (partCount);
		setParticleSizes (partCount);
	}

	void setParticleColors(int partCount) {
		for (int idx = 0; idx < partCount; idx ++) {
			if(particles[idx].startColor.a == 0f) {
				color.Mutate ();
				particles[idx].startColor = new Color (color.R, color.G, color.B, color.A);
			}
		}

		ps.SetParticles (particles, partCount);
	}

	void setParticleSizes(int partCount) {
		for (int idx = 0; idx < particles.Length; idx ++) {
			particles[idx].startSize = 2f;
		}
		ps.SetParticles (particles, partCount);
	}

	//TODO
	void setParticleLifetimes(int partCount) {
		ps.SetParticles (particles, partCount);
	}
}