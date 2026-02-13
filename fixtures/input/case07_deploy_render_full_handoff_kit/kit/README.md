# Casefile: Deploy Render Handoff Kit (Full Config)

This kit demonstrates the **repo-to-repo handoff** from **Repo B** (`proof-first-deploy-gcp`) into **Repo C** (`proof-first-casefiles`).

## Scenario
A client wants a deploy plan they can review and sign off on **before** anything touches GCP.

You hand them:
- the `config.yaml` you intend to deploy
- the deterministic render outputs (`*_manifest.json`)
- the `manifest.sha256` witness file

## How to reproduce (Repo B)
From the `proof-first-deploy-gcp` repo (Book 2 snapshot tag, e.g. `book2-v1`):

```bash
go run ./cmd/pfdeploy render --config ./fixtures/input/case02_render_full/config.yaml --out ./out/render
```

Then compare `./out/render/*` to this kit’s `inputs/render/*`.

## Why this matters
These artifacts are:
- diffable (PR-friendly)
- hashable (audit-friendly)
- safe to share (no secrets)
