# Names-only review paths (Book 4 PR discipline)

A reviewer should be able to evaluate the PR without reading the whole repo.

## Inspect the kit inputs

- `fixtures/input/case10_scope_contract_smoke/kit/kit.json`
- `fixtures/input/case10_scope_contract_smoke/kit/README.md`
- `fixtures/input/case10_scope_contract_smoke/kit/inputs/scope_contract.json`
- `fixtures/input/case10_scope_contract_smoke/kit/inputs/acceptance.md`
- `fixtures/input/case10_scope_contract_smoke/kit/inputs/review_paths.md`
- `fixtures/input/case10_scope_contract_smoke/kit/inputs/fixture_map.md`

## Inspect the receipts

- `fixtures/expected/case10_scope_contract_smoke/kit_index.json`
- `fixtures/expected/case10_scope_contract_smoke/manifest.sha256`

## Run the proof gate

```bash
make verify
```
