package main

import (
	"path/filepath"
	"testing"
)

func TestProcessKitCase01(t *testing.T) {
	kit := filepath.Join("..", "..", "fixtures", "input", "case01_bank_recon_kit", "kit")
	out, err := ProcessKit(kit)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if out.Index.Title == "" || len(out.Index.Files) == 0 {
		t.Fatalf("bad index")
	}
}

func TestProcessKitExpectedFailMissingInput(t *testing.T) {
	kit := filepath.Join("..", "..", "fixtures", "input", "case03_missing_input_expected_fail", "kit")
	_, err := ProcessKit(kit)
	if err == nil {
		t.Fatalf("expected error")
	}
}
