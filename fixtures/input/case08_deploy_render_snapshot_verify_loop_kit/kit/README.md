# Casefile: Deploy Handoff Loop Kit (Render + Snapshot + Verify)

This kit closes the **end-to-end handoff loop** across Book 2:

**Repo B** (`proof-first-deploy-gcp`) produces deterministic artifacts, then later verifies them against a live snapshot.

## What’s inside
Render receipts (what you *intended* to deploy):
- `inputs/render/deploy_manifest.json`
- `inputs/render/trigger_manifest.json`
- `inputs/render/iam_manifest.json`
- `inputs/render/manifest.sha256`

Snapshot + verify (what is *actually* running):
- `inputs/snapshot/gcloud_service.json`
- `inputs/verify/verify_report.json`

## How to reproduce (Repo B)
From the `proof-first-deploy-gcp` repo (Book 2 snapshot tag, e.g. `book2-v1`):

```bash
# 1) Render (plan)
go run ./cmd/pfdeploy render --config ./fixtures/input/case05_verify_matches/config.yaml --out ./out/render

# 2) Verify (compare snapshot vs config)
go run ./cmd/pfdeploy verify --config ./fixtures/input/case05_verify_matches/config.yaml --snapshot ./fixtures/input/case05_verify_matches/gcloud_service.json --out ./out/verify
```

Then compare:
- `./out/render/*` to `inputs/render/*`
- `./out/verify/verify_report.json` to `inputs/verify/verify_report.json`

## Why this matters
This is the “trust story” Book 2 is selling:
- deterministic plan → deterministic receipts → snapshot → deterministic verification
