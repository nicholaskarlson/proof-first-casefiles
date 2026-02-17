# case10_scope_contract_smoke — Scope Contract (Smoke)

This kit is the smallest “consulting-grade receipt” in the Book 4 ladder.

It answers three questions *before* we build anything:

1) **What is in scope?**
2) **How will a reviewer verify it (names-only paths)?**
3) **What does “done” mean (proof gate + deterministic artifacts)?**

## How this fixture is used

The `pfcasefiles demo` command renders this kit into an output folder and compares
it byte-for-byte against the goldens in:

- `fixtures/expected/case10_scope_contract_smoke/`

If the kit changes, the receipt changes — and the proof gate catches it.

## Names-only review paths

Reviewers should inspect only these paths:

- `fixtures/input/case10_scope_contract_smoke/kit/kit.json`
- `fixtures/input/case10_scope_contract_smoke/kit/inputs/scope_contract.json`
- `fixtures/input/case10_scope_contract_smoke/kit/inputs/acceptance.md`
- `fixtures/input/case10_scope_contract_smoke/kit/inputs/review_paths.md`
- `fixtures/input/case10_scope_contract_smoke/kit/inputs/fixture_map.md`

And then confirm the goldens match:

- `fixtures/expected/case10_scope_contract_smoke/kit_index.json`
- `fixtures/expected/case10_scope_contract_smoke/manifest.sha256`

## Proof gate

One command:

```bash
make verify
```
