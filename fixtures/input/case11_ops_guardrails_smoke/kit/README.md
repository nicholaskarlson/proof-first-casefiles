# Casefile: Ops Guardrails Smoke (Book 4 Ch12)

This kit is the **operations contract** you hand to a client (or your future self) so the system has
**guardrails you can verify** — not “we’ll watch the dashboard.”

## What this kit locks
- **Cost ceiling**: budgets and “stop-the-bleed” triggers.
- **Blast radius limits**: concurrency caps, max object sizes, allowed prefixes.
- **Timeouts + retries**: deterministic policies (no surprise infinite loops).
- **Alert routing**: who gets paged, what evidence to attach, what *not* to do.
- **Names-only review**: reviewers start from `kit_index.json` + `manifest.sha256` and inspect only named artifacts.

## How to reproduce outputs
From the repo root:

```bash
go run ./cmd/pfcasefiles render --kit ./fixtures/input/case11_ops_guardrails_smoke/kit --out ./out/ch12
```

Then compare:

```bash
git diff --no-index ./fixtures/expected/case11_ops_guardrails_smoke ./out/ch12
```

## Why this matters
Guardrails are part of the **receipt chain**. When something goes wrong, you want:
- a deterministic policy you can point to, and
- evidence that you followed it.
