using System.Collections;
using System.Collections.Generic;
using UnityEngine;

//https://docs.unity3d.com/ScriptReference/Application.CaptureScreenshot.html
public class ScreenCapure : MonoBehaviour {
	void OnMouseDown() {
		ScreenCapture.CaptureScreenshot("flowFields.png");
		Debug.Log ("Captured");
	}
}
