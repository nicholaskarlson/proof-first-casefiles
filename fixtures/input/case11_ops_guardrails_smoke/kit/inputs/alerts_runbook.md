# Alerts runbook (ops guardrails)

When a guardrail triggers, the goal is **evidence-first triage**.

## Always attach these receipts
- `plan_manifest.json` + `manifest.sha256` (what we intended to do)
- `batch_report.json` + `manifest.sha256` (what we actually did)
- `verify_report.json` (what matched) or `error.txt` (what failed honestly)
- `drift_report.json` (what changed between runs)

## Common triggers
- **Budget threshold** (daily/monthly): pause new runs; open a change review with the receipts above.
- **Retry ceiling** exceeded: treat as a hard failure; do not “just re-run” without a drift diff.
- **Timeout**: capture the run metadata; verify manifests before retrying.
- **Unexpected object key / prefix**: treat as suspicious input; block and escalate.

## What not to do
- Don’t delete evidence.
- Don’t “fix forward” by editing outputs in place.
- Don’t rerun with different config unless you also produce a new plan manifest.
