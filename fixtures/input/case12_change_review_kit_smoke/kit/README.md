# Casefile: Change Review Kit Smoke (Book 4 Ch13)

This kit is the smallest **consulting-grade change review receipt**.

It answers the question a reviewer actually has:

> **What changed, why, what could break, and what evidence proves it?**

The trick is that the answers are **deterministic artifacts**, not meeting notes.

## What this kit locks
- **Change intent**: the one-paragraph “why” and the bounded scope of the change.
- **Review workflow**: names-only paths a reviewer can follow without spelunking.
- **Risk + rollback**: what could go wrong and the “get out safely” plan.
- **Acceptance**: what “done” means in testable terms.
- **Fixture mapping**: which fixtures are the proof for this chapter milestone.

## How to reproduce outputs
From the repo root:

```bash
go run ./cmd/pfcasefiles render --kit ./fixtures/input/case12_change_review_kit_smoke/kit --out ./out/ch13
```

Then compare the rendered receipt against goldens:

```bash
git diff --no-index ./fixtures/expected/case12_change_review_kit_smoke ./out/ch13
```

## Names-only review paths
Start from the receipts:

- `fixtures/expected/case12_change_review_kit_smoke/kit_index.json`
- `fixtures/expected/case12_change_review_kit_smoke/manifest.sha256`

Then inspect only named kit inputs:

- `fixtures/input/case12_change_review_kit_smoke/kit/kit.json`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/change_request.json`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/acceptance.md`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/risk_rollback.md`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/review_paths.md`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/fixture_map.md`

## Proof gate

One command:

```bash
make verify
```
