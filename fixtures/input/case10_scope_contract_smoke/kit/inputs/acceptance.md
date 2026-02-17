# Acceptance — Scope Contract Smoke

This fixture is *done* when:

- The kit is valid (`kit.json` has required fields and listed inputs exist).
- `pfcasefiles demo --out ./out` produces deterministic receipts for this kit.
- The receipts match goldens byte-for-byte:

  - `fixtures/expected/case10_scope_contract_smoke/kit_index.json`
  - `fixtures/expected/case10_scope_contract_smoke/manifest.sha256`

- The proof gate passes:

```bash
make verify
```

## Boundaries (what this fixture is *not*)

- It does not validate cloud state.
- It does not prove deploy or backfill correctness.
- It does not include secrets, tokens, or credentials.
