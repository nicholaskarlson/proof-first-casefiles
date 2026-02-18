# Fixture map (Book 4)

This kit is a **casefile wrapper** around receipts produced by the Book 4 workflow repos.

## Upstream receipts this ops kit expects (names only)

From `proof-first-deploy-gcp`:
- `fixtures/expected/case01_deploy_render_smoke/manifest.sha256`
- `fixtures/expected/case02_snapshot_smoke/manifest.sha256`
- `fixtures/expected/case03_verify_matches/verify_report.json`
- `fixtures/expected/case04_verify_mismatch_expected_fail/error.txt` (expected-fail)

From `proof-first-backfills-gcp`:
- `fixtures/expected/case02_cloud_plan_consulting_smoke/cloud/cloud_plan.json`
- `fixtures/expected/case03_apply_smoke_consulting/apply/batch_report.json`
- `fixtures/expected/case04_verify_smoke_consulting/verify/verify_report.json`
- `fixtures/expected/case05_verify_apply_manifest_mismatch_expected_fail/error.txt` (expected-fail)
- `fixtures/expected/case06_pack_smoke_consulting/pack/pack_manifest.json`
- `fixtures/expected/case07_drift_one_change/diff/drift_report.json`

## What this chapter proves
Ops is not “extra.” It’s where deterministic systems survive contact with reality:
guardrails become receipts you can review, test, and hand off.
