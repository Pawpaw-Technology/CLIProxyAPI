package executor

import "testing"

func TestGetVertexAction_ClaudeUsesRawPredict(t *testing.T) {
	if got := getVertexAction("claude-sonnet-4-5-20250929", false); got != "rawPredict" {
		t.Fatalf("expected rawPredict, got %s", got)
	}
	if got := getVertexAction("claude-sonnet-4-5-20250929", true); got != "streamRawPredict" {
		t.Fatalf("expected streamRawPredict, got %s", got)
	}
}

func TestGetVertexAction_GeminiAndImagen(t *testing.T) {
	if got := getVertexAction("gemini-2.5-pro", false); got != "generateContent" {
		t.Fatalf("expected generateContent, got %s", got)
	}
	if got := getVertexAction("gemini-2.5-pro", true); got != "streamGenerateContent" {
		t.Fatalf("expected streamGenerateContent, got %s", got)
	}
	if got := getVertexAction("imagen-4.0-generate-preview-06-06", false); got != "predict" {
		t.Fatalf("expected predict, got %s", got)
	}
}

func TestVertexPublisherAndFormat(t *testing.T) {
	if got := getVertexPublisher("claude-sonnet-4-5-20250929"); got != "anthropic" {
		t.Fatalf("expected anthropic publisher, got %s", got)
	}
	if got := getVertexPublisher("gemini-2.5-pro"); got != "google" {
		t.Fatalf("expected google publisher, got %s", got)
	}
	if got := vertexTargetFormat("claude-sonnet-4-5-20250929"); got != "claude" {
		t.Fatalf("expected claude format, got %s", got)
	}
	if got := vertexTargetFormat("gemini-2.5-pro"); got != "gemini" {
		t.Fatalf("expected gemini format, got %s", got)
	}
}

func TestVertexModelURL_UsesPublisherByModel(t *testing.T) {
	urlClaude := vertexServiceAccountModelURL("global", "proj-x", "claude-sonnet-4-5-20250929", "rawPredict")
	if urlClaude != "https://aiplatform.googleapis.com/v1/projects/proj-x/locations/global/publishers/anthropic/models/claude-sonnet-4-5-20250929:rawPredict" {
		t.Fatalf("unexpected claude url: %s", urlClaude)
	}

	urlGemini := vertexServiceAccountModelURL("global", "proj-x", "gemini-2.5-pro", "generateContent")
	if urlGemini != "https://aiplatform.googleapis.com/v1/projects/proj-x/locations/global/publishers/google/models/gemini-2.5-pro:generateContent" {
		t.Fatalf("unexpected gemini url: %s", urlGemini)
	}

	urlAPIKeyClaude := vertexAPIKeyModelURL("https://aiplatform.googleapis.com", "claude-sonnet-4-5-20250929", "rawPredict")
	if urlAPIKeyClaude != "https://aiplatform.googleapis.com/v1/publishers/anthropic/models/claude-sonnet-4-5-20250929:rawPredict" {
		t.Fatalf("unexpected api-key claude url: %s", urlAPIKeyClaude)
	}
}
