# Review paths (names-only)

Start from the rendered receipts:

1) `kit_index.json`  
   - Confirms the kit identity (`title`, `category`), file list, and byte sizes.

2) `manifest.sha256`  
   - A checksum witness for every file in the kit (sorted by sha256).

Then inspect only the named contract artifacts:

- `inputs/ops_guardrails.json`  
  Guardrail contract: cost, limits, timeouts, retries, alert routing.

- `inputs/alerts_runbook.md`  
  What to do when the guardrails fire (and what evidence to attach).

- `inputs/acceptance.md`  
  The minimum bar for “safe to run” in consulting terms.

- `inputs/fixture_map.md`  
  Which upstream receipts this kit expects to exist in the handoff chain.
