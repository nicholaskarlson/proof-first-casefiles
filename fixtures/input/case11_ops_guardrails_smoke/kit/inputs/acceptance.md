# Acceptance (Ops Guardrails)

A reviewer should be able to answer “are we safe to run this?” by reading **only** the named artifacts.

## Must-haves
- A written **budget ceiling** (daily + monthly) and a “stop-the-bleed” threshold.
- A **concurrency cap** and a maximum per-run batch size.
- A deterministic **timeout + retry** policy (no infinite retries; no silent backoff drift).
- A clear **alert path** (who, when, with what evidence).
- A names-only review path that points to the exact receipts a run must produce.

## Done means done
This kit renders deterministically and passes the proof gate:

```bash
make verify
```
