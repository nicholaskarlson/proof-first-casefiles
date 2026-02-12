# Conventions

## Line endings
- LF only (enforced via `.gitattributes`).
- This repo also *rejects* CRLF in text-like kit files to keep casefiles portable.

## Determinism
Artifacts produced by this repo are deterministic:
- file lists are sorted
- JSON output is stable (`json.MarshalIndent`) and ends with a trailing newline
- `manifest.sha256` uses sorted paths and sha256 over raw bytes

## Expected-fail
Expected-fail casefiles emit **only**:
- `error.txt` (must end with a newline)
