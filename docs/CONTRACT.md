# Contract (Repo C)

Input (per case):
- `kit/README.md` (one-page brief)
- `kit/kit.json` (metadata + rules)
- `kit/inputs/...` (CSV/JSON/etc.)

Outputs:
- success:
  - `kit_index.json`
  - `manifest.sha256`
- expected-fail:
  - `error.txt` only

Rules enforced:
1) Required input files listed in `kit.json` must exist.
2) Text-like files (`.md .json .csv .txt .yaml .yml`) must be LF-only (no CRLF).
3) No secrets in env keys inside `kit.json` (`SECRET`, `TOKEN`, `PASSWORD`, `KEY`).
