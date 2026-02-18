# Risk + Rollback

## Risk surface (what could break)
- **Receipt drift**: a harmless-looking edit changes bytes → manifests change → proof gate fails.
- **Hidden scope creep**: new required artifacts added without being named in review paths.
- **Process failure**: reviewers forced into “explore mode” because the kit doesn't say where to look.

## Rollback plan
If this change causes unexpected churn:
1) Revert the PR commit that introduced `case12_change_review_kit_smoke`.
2) Re-run the proof gate:

```bash
make verify
```

3) Re-open the change with a tighter scope and updated kit inputs (no silent edits).

## What we do NOT do
- Do not “fix” a failing proof gate by editing goldens without a documented reason.
- Do not bypass manifests or accept non-deterministic outputs.
