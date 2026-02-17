# Delivery kit (smoke)

This folder is an example **offline-review** kit.

A reviewer should be able to validate integrity without cloud access:

## 1) Intake (casefile)

```bash
(cd ./00-intake/kit && sha256sum -c ../manifest.sha256)
```

Open:
- `00-intake/kit_index.json`

## 2) Deploy render receipt

```bash
(cd ./10-deploy/render && sha256sum -c manifest.sha256)
```

Open:
- `10-deploy/verify/verify_report.json`

## 3) Backfill packs

```bash
(cd ./20-backfill/pack_A && sha256sum -c manifest.sha256)
(cd ./21-backfill/pack_B && sha256sum -c manifest.sha256)
```

## 4) Drift report

```bash
(cd ./30-drift && sha256sum -c manifest.sha256)
```

Open:
- `30-drift/drift_report.json`
