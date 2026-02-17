# Casefile: Delivery Kit Smoke (Intake + Deploy + Backfill + Drift)

This kit is a **shipping-shaped** folder: it contains an example `delivery_kit/` tree you can hand to a reviewer.

The point is not “more automation.” The point is **bounded review**:
- every deterministic evidence folder includes a `manifest.sha256`
- drift is a single `drift_report.json` that names the exact paths that changed

## What’s inside (names only)
The example delivery kit lives here:

- `inputs/delivery_kit/README.md`
- `inputs/delivery_kit/00-intake/` (casefile kit + evidence)
- `inputs/delivery_kit/10-deploy/` (render + snapshot + verify)
- `inputs/delivery_kit/20-backfill/pack_A/` (pack output)
- `inputs/delivery_kit/21-backfill/pack_B/`
- `inputs/delivery_kit/30-drift/` (drift report)

## How to use this fixture
From the casefiles repo:

```bash
# render an index + manifest for the entire delivery kit casefile
go run ./cmd/pfcasefiles render \
  --kit ./fixtures/input/case09_delivery_kit_smoke/kit \
  --out ./out/case09

# optional: a tiny verification summary (counts + bytes)
go run ./cmd/pfcasefiles verify \
  --kit ./fixtures/input/case09_delivery_kit_smoke/kit \
  --out ./out/case09_verify
```

Then open `./out/case09/kit_index.json` and review the delivered paths.
