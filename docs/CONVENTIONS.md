# Conventions

These conventions exist so the repo can *prove* its outputs: deterministic artifacts + fixtures + goldens + a verification gate.

## Line endings
- **LF only** (enforced via `.gitattributes`).
- This repo also *rejects* CRLF in text-like kit files to keep casefiles portable.

## Determinism
Artifacts produced by this repo are deterministic:
- file lists are sorted
- JSON output is stable (`json.MarshalIndent`) and ends with a trailing newline
- `kit_index.json` lists files in sorted path order
- `manifest.sha256` hashes raw file bytes; each line is `SHA256␠␠path` and lines are sorted lexicographically (sha-first)
- no timestamps, UUIDs, random IDs, or host-specific paths embedded into artifacts

## Atomic writes
Artifacts are written via temp file → rename so partial files never appear.

## Output directory
- Commands that write artifacts (`render`, `verify`, `demo`) **clear the `--out` directory first**.
- A safety guard refuses unsafe paths (e.g., `.`, `..`, `/`, Windows volume roots).

## Expected-fail
Expected-fail casefiles emit **only**:
- `error.txt` (must end with a newline)

The proof gate (`make verify`) recomputes outputs and compares them byte-for-byte to `fixtures/expected/**`.
