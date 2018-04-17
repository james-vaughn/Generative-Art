using System.Collections;
using System.IO;
using System.Collections.Generic;
using UnityEngine;

//https://docs.unity3d.com/ScriptReference/Application.CaptureScreenshot.html
public class ScreenCapure : MonoBehaviour {

	private bool m_screenShotLock = false;

	private void LateUpdate()
	{
		if (Input.GetKeyDown(KeyCode.S) && !m_screenShotLock)
		{
			m_screenShotLock = true;
			StartCoroutine(TakeScreenShotCo());
		}
	}

	private IEnumerator TakeScreenShotCo()
	{
		yield return new WaitForEndOfFrame();

		var directory = new DirectoryInfo(Application.dataPath);
		var path = Path.Combine(directory.Parent.FullName, "flowFields.png");
		Debug.Log("Taking screenshot to " + path);
		ScreenCapture.CaptureScreenshot(path);
		m_screenShotLock = false;
	}
}
