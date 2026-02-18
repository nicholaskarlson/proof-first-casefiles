# Acceptance (Change Review Kit Smoke)

This fixture is accepted when:

1) `go run ./cmd/pfcasefiles render --kit ... --out ...` produces deterministic receipts.
2) The rendered receipts match the goldens byte-for-byte:

   - `fixtures/expected/case12_change_review_kit_smoke/kit_index.json`
   - `fixtures/expected/case12_change_review_kit_smoke/manifest.sha256`

3) The proof gate passes:

```bash
make verify
```

Non-negotiables:
- LF line endings for text-like files.
- No secrets / env keys / credentials.
- Names-only review paths (no “go browse the repo until it feels right”).
