package client

import "testing"

// TestSignerByteEquivalence locks the recovered signer against a captured
// real request. Inputs and expected output extracted from
// recon/traffic/real/0001_GET_project_v3_projects_PROJECT_nMnc4vToqXfIZDQZfHMYMf.json
// taken on 2026-05-14 against openapi.samsungsdscloud.com.
//
// If this test ever fails, the canonical-request composition or one of the
// header inputs has drifted. Re-capture a flow and bisect.
func TestSignerByteEquivalence(t *testing.T) {
	const (
		method    = "GET"
		fullURL   = "https://localhost:8443/project/v3/projects/PROJECT-nMnc4vToqXfIZDQZfHMYMf"
		timestamp = "1778728850883"
		accessKey = "9136/5LY5XCQcurvagdP"
		secretKey = "cnZhZ2RQc29oNE4zeW5YOEtRdHRUN0UvSVA4dFE9"
		projectId = "PROJECT-nMnc4vToqXfIZDQZfHMYMf"
		want      = "eOWX9f91sSK53HadviGWujYdoTTmENunE7brrV6x5r0="
	)
	got := MakeHmacSignature(method, fullURL, timestamp, accessKey, secretKey, projectId, "OpenApi")
	if got != want {
		t.Fatalf("signature mismatch:\n  got  %q\n  want %q", got, want)
	}
}
