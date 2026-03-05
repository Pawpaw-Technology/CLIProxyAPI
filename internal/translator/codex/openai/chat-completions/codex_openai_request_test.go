package chat_completions

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestConvertOpenAIRequestToCodex_PreservesExplicitSummaryAndInclude(t *testing.T) {
	input := []byte(`{
		"model":"gpt-5.3-codex",
		"messages":[{"role":"user","content":"hi"}],
		"reasoning_effort":"high",
		"reasoning":{"summary":"detailed"},
		"include":["reasoning.encrypted_content","reasoning.summary"]
	}`)

	output := ConvertOpenAIRequestToCodex("gpt-5.3-codex", input, false)
	out := string(output)

	if got := gjson.Get(out, "reasoning.effort").String(); got != "high" {
		t.Fatalf("reasoning.effort mismatch: got %q, want %q", got, "high")
	}

	if got := gjson.Get(out, "reasoning.summary").String(); got != "detailed" {
		t.Fatalf("reasoning.summary mismatch: got %q, want %q", got, "detailed")
	}

	include := gjson.Get(out, "include")
	if !include.Exists() || !include.IsArray() {
		t.Fatalf("include should be present array, got: %s", include.Raw)
	}
	if got := include.Array()[0].String(); got != "reasoning.encrypted_content" {
		t.Fatalf("include[0] mismatch: got %q, want %q", got, "reasoning.encrypted_content")
	}
	if got := include.Array()[1].String(); got != "reasoning.summary" {
		t.Fatalf("include[1] mismatch: got %q, want %q", got, "reasoning.summary")
	}
}

func TestConvertOpenAIRequestToCodex_FallsBackToDefaultsWhenMissing(t *testing.T) {
	input := []byte(`{
		"model":"gpt-5.3-codex",
		"messages":[{"role":"user","content":"hi"}],
		"reasoning_effort":"medium"
	}`)

	output := ConvertOpenAIRequestToCodex("gpt-5.3-codex", input, false)
	out := string(output)

	if got := gjson.Get(out, "reasoning.summary").String(); got != "auto" {
		t.Fatalf("reasoning.summary default mismatch: got %q, want %q", got, "auto")
	}

	include := gjson.Get(out, "include")
	if !include.Exists() || !include.IsArray() {
		t.Fatalf("default include should be present array, got: %s", include.Raw)
	}
	if got := include.Array()[0].String(); got != "reasoning.encrypted_content" {
		t.Fatalf("default include[0] mismatch: got %q, want %q", got, "reasoning.encrypted_content")
	}
}
