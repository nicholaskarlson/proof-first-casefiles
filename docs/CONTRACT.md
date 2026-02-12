# Contract (Repo C)

This repo converts a casefile “kit” into deterministic evidence artifacts.

## Input shape (per case)

A case lives under:

- `fixtures/input/<case>/kit/`

Required kit files:

- `kit/README.md` — one-page brief
- `kit/kit.json` — metadata + rules
- `kit/inputs/**` — input files (CSV/JSON/etc.)

Optional:

- `kit/notes/**` — supporting materials

## Outputs

### Success
- `kit_index.json`
- `manifest.sha256`

### Expected-fail (contract violation)
- `error.txt` only

### Verify (optional report)
- `verify_report.json` (success) OR `error.txt` (contract violation)

All text outputs end with a trailing newline.

## Rules enforced

1) **Required inputs exist**  
   Files listed in `kit.json` must exist under the kit directory.

2) **LF-only text-like files**  
   Text-like extensions are treated as “must be LF”:
   `.md .json .csv .txt .yaml .yml`  
   CRLF is rejected (the error is written to `error.txt`).

3) **No secrets in env keys**  
   `kit.json` env keys are rejected if they contain (case-insensitive):
   `SECRET`, `TOKEN`, `PASSWORD`, `KEY`.

4) **Deterministic artifacts**  
   - file lists are sorted
   - JSON output is stable (`json.MarshalIndent`) + trailing newline
   - `manifest.sha256` uses sorted paths and sha256 over raw bytes

## CLI semantics (important for automation)

- `pfcasefiles render --kit ... --out ...`
  - clears the out dir before writing
  - writes `error.txt` (and returns exit code 0) on expected-fail

- `pfcasefiles demo --out ...`
  - clears the out dir
  - recomputes outputs for every fixture under `fixtures/input/**`
  - diffs against `fixtures/expected/**` and fails on mismatch

- `pfcasefiles verify --kit ... --out ...`
  - clears the out dir
  - writes `verify_report.json` on success, else `error.txt`
