package registry

import "testing"

func TestGetGeminiVertexModels_IncludesClaudePartnerModels(t *testing.T) {
	models := GetGeminiVertexModels()
	if len(models) == 0 {
		t.Fatal("expected vertex models, got empty list")
	}

	required := map[string]struct {
		ownedBy string
		typ     string
	}{
		"claude-sonnet-4-5@20250929": {ownedBy: "anthropic", typ: "claude"},
		"claude-opus-4-5@20251101":   {ownedBy: "anthropic", typ: "claude"},
	}

	index := make(map[string]*ModelInfo, len(models))
	for _, m := range models {
		if m == nil || m.ID == "" {
			continue
		}
		index[m.ID] = m
	}

	for modelID, expected := range required {
		model, ok := index[modelID]
		if !ok {
			t.Fatalf("expected vertex model %q to exist", modelID)
		}
		if model.OwnedBy != expected.ownedBy {
			t.Fatalf("model %q owned_by mismatch: got %q want %q", modelID, model.OwnedBy, expected.ownedBy)
		}
		if model.Type != expected.typ {
			t.Fatalf("model %q type mismatch: got %q want %q", modelID, model.Type, expected.typ)
		}
	}
}
