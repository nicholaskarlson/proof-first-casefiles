# proof-first-casefiles

Repo C: **Freelancer engagement kits** ("casefiles") with deterministic, verifiable handoffs.

Each casefile is a small, realistic data package:
- a one-page brief (README)
- input files (CSV/JSON)
- `kit.json` metadata (what the kit contains / rules)
- deterministic evidence artifacts produced by this repo:
  - `kit_index.json`
  - `manifest.sha256`
  - or `error.txt` for expected-fail kits

Quickstart:
- `make verify`
- `go run ./cmd/pfcasefiles demo --out ./out`

## Casefile rules
- LF-only text files (CRLF is rejected)
- No secrets in `kit.json` env keys (SECRET/TOKEN/PASSWORD/KEY)
- Deterministic ordering (sorted paths, stable JSON formatting)

See:
- `docs/CONTRACT.md`
- `docs/CONVENTIONS.md`
- `docs/HANDOFF.md`
