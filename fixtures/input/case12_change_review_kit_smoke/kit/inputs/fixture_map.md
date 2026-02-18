# Fixture map (what proves this chapter milestone)

This kit is the **proof artifact** for Book 4 Chapter 13.

Inputs (kit):
- `fixtures/input/case12_change_review_kit_smoke/kit/kit.json`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/*`

Outputs (receipts):
- `fixtures/expected/case12_change_review_kit_smoke/kit_index.json`
- `fixtures/expected/case12_change_review_kit_smoke/manifest.sha256`

Reviewer flow:
1) Read `kit_index.json` (file list + sha256 + bytes).
2) Confirm `manifest.sha256` matches the same file set (names-only).
3) Inspect only the named inputs.
