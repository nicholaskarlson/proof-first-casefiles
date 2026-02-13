# Handoff (Repo C)

A **casefile** is meant to be attachable to a client handoff, audit pack, or course module.

## What you hand off

For a successful kit, hand off:

- the kit itself (`kit/`):
  - `README.md` (one-page brief)
  - `kit.json` (rules + inventory)
  - `inputs/**` (data)
  - optional `notes/**`
- the evidence artifacts produced by this repo:
  - `kit_index.json`
  - `manifest.sha256`
  - optional: `verify_report.json`

For expected-fail kits, hand off:

- the kit itself (`kit/`)
- `error.txt` (the contract violation message)

## Recommended workflow (human-friendly)

1) Render evidence artifacts (clears the output dir first):

```bash
go run ./cmd/pfcasefiles render --kit <kit_dir> --out <out_dir>
```

2) (Optional) Generate a verification summary for the handoff bundle:

```bash
go run ./cmd/pfcasefiles verify --kit <kit_dir> --out <out_dir>
```

3) Attach:
- the kit folder
- the evidence outputs (`kit_index.json`, `manifest.sha256`, and optionally `verify_report.json`)

## Guardrails

- This repo refuses “secrets” in `kit.json` env keys (eg `TOKEN`, `PASSWORD`, `SECRET`, `KEY`).
- Text-like files must be LF-only; CRLF is rejected to keep casefiles portable.