# proof-first-casefiles

Repo C (Book 2): **Freelancer engagement kits** (“casefiles”) with deterministic, verifiable handoffs.

A casefile is a small, realistic data package plus a **proof gate** that recomputes evidence artifacts and compares them byte-for-byte to goldens.

**Go baseline:** 1.22.x (CI witnesses ubuntu/macos/windows on 1.22.x, plus ubuntu “stable”).

## Where this fits (Book 2)

This repo is designed to be used alongside the other Book 2 repos:

- Anchor: `finance-pipeline-gcp` — deployable drop-folder workflow (trigger → run → artifacts → markers)
- Repo A: `proof-first-event-contracts` — event parsing contract + fixtures/goldens + expected-fail
- Repo B: `proof-first-deploy-gcp` — deterministic deploy evidence (render + verify) + fixtures/goldens
- Repo C: `proof-first-casefiles` — engagement kits you can hand to a client (or use in teaching)

## Quickstart

Run the proof gate:

```bash
gofmt -w cmd internal
make verify
go test -count=1 ./...
```

Run the deterministic fixture demo (recomputes outputs and diffs against `fixtures/expected/**`):

```bash
go run ./cmd/pfcasefiles demo --out ./out
```

Render evidence artifacts for a single kit:

```bash
go run ./cmd/pfcasefiles render --kit fixtures/input/case01_bank_recon_kit/kit --out ./out/render1
```

Generate a lightweight verification report (for handoff):

```bash
go run ./cmd/pfcasefiles verify --kit fixtures/input/case01_bank_recon_kit/kit --out ./out/verify1
```

## Fixture layout

Inputs live under:

- `fixtures/input/<case>/kit/README.md` (one-page brief)
- `fixtures/input/<case>/kit/kit.json` (metadata + rules)
- `fixtures/input/<case>/kit/inputs/**` (CSV/JSON/etc.)
- optional: `fixtures/input/<case>/kit/notes/**` (supporting text)

Goldens live under:

- `fixtures/expected/<case>/kit_index.json` + `manifest.sha256` **OR**
- `fixtures/expected/<case>/error.txt` (expected-fail cases)

## Output artifacts

On success, this repo emits:

- `kit_index.json` — normalized inventory + per-file sha256 + totals (stable JSON + trailing newline)
- `manifest.sha256` — sorted `sha256  relative/path` lines over the kit tree (trailing newline)

For expected-fail (contract violations), the output is:

- `error.txt` only (must end with a newline)

`verify` emits:

- `verify_report.json` (success) **or** `error.txt` (contract violation)

## Rules enforced (book-friendly, audit-friendly)

- **LF-only** text-like files (`.md .json .csv .txt .yaml .yml`); CRLF is rejected.
  - One fixture (`case05_crlf_text_expected_fail`) intentionally contains CRLF bytes to prove the guard.
- **No secrets** in `kit.json` env keys (case-insensitive match on: `SECRET`, `TOKEN`, `PASSWORD`, `KEY`).
- **Deterministic ordering**: sorted file lists and stable JSON formatting.
- **No entropy** in artifacts: no timestamps, UUIDs, random IDs, or host-specific paths.

## Add a new casefile (the “Book 2” pattern)

1) Create `fixtures/input/caseXX_some_name/kit/**` (brief + kit.json + inputs).
2) Generate outputs into a scratch out dir:

```bash
go run ./cmd/pfcasefiles render --kit fixtures/input/caseXX_some_name/kit --out ./out/tmp
```

### Included fixture kits

- `case01_bank_recon_kit` — basic left/right CSV recon kit.
- `case02_audit_pack_kit` — audit-pack bundle (postings + issues).
- `case03_missing_input_expected_fail` — expected-fail: missing required input.
- `case04_secret_envkey_expected_fail` — expected-fail: forbidden env key name (secrets).
- `case05_crlf_text_expected_fail` — expected-fail: CRLF bytes in a text-like file.
- `case06_multi_file_ordering_kit` — ordering demo across multiple files/folders.
- `case07_deploy_render_full_handoff_kit` — **Repo B handoff**: render receipts + sha256 witness.
- `case08_deploy_render_snapshot_verify_loop_kit` — **end-to-end loop**: render + snapshot + verify artifacts.

3) Copy outputs into `fixtures/expected/caseXX_some_name/` and commit.
4) Run the gate:

```bash
make verify
```

See also:
- `docs/CONTRACT.md`
- `docs/CONVENTIONS.md`
- `docs/HANDOFF.md`


The book references **tags** (e.g. `book2-v1`) across all four repos, not moving `main` branches.
