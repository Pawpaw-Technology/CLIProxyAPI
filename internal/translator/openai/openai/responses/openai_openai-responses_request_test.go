package responses

import (
	"testing"

	"github.com/tidwall/gjson"
)

func TestConvertOpenAIResponsesRequestToOpenAIChatCompletions_MapsReasoningSummaryAndInclude(t *testing.T) {
	input := []byte(`{
		"model":"gpt-5.3-codex",
		"input":[{"role":"user","content":"hi"}],
		"reasoning":{"effort":"high","summary":"detailed"},
		"include":["reasoning.encrypted_content","reasoning.summary"]
	}`)

	output := ConvertOpenAIResponsesRequestToOpenAIChatCompletions("gpt-5.3-codex", input, true)
	out := string(output)

	if got := gjson.Get(out, "reasoning_effort").String(); got != "high" {
		t.Fatalf("reasoning_effort mismatch: got %q, want %q", got, "high")
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
