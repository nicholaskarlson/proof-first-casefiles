# Review paths (names-only)

Start at:
- `kit_index.json`
- `manifest.sha256`

Then open only these named artifacts:

## Graduation index
- `inputs/graduation_kit/README.md`
- `inputs/graduation_kit/witness_matrix.md`

## Contract
- `inputs/graduation_kit/00_contract/scope_contract.json`

## Strict event decode (input contract)
- `inputs/graduation_kit/10_event_contracts/event.json`
- `inputs/graduation_kit/10_event_contracts/decision.json`
- `inputs/graduation_kit/10_event_contracts/object_ref.json`

## Plan + execution bundle + drift
- `inputs/graduation_kit/20_backfills/cloud/cloud_plan.json`
- `inputs/graduation_kit/21_backfills_pack/pack/pack_manifest.json`
- `inputs/graduation_kit/22_drift/diff/drift_report.json`

## Deploy evidence (render + snapshot + verify)
- `inputs/graduation_kit/30_deploy_render/manifest.sha256`
- `inputs/graduation_kit/31_deploy_snapshot/manifest.sha256`
- `inputs/graduation_kit/32_deploy_verify/verify_report.json`

## Ops + change review
- `inputs/graduation_kit/60_ops/ops_guardrails.json`
- `inputs/graduation_kit/60_ops/alerts_runbook.md`
- `inputs/graduation_kit/70_change_review/change_request.json`
- `inputs/graduation_kit/70_change_review/risk_rollback.md`

Lineage:
- `inputs/fixture_map.md`
