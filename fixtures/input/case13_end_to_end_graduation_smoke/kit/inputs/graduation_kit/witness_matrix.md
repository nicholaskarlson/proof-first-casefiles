# Witness matrix (what to run)

This graduation kit packages receipts produced by proof gates in sibling repos:

- `proof-first-event-contracts`: `make verify`
- `proof-first-backfills-gcp`: `make verify`
- `proof-first-deploy-gcp`: `make verify`
- `proof-first-casefiles`: `make verify`

If the gates are green, a reviewer can treat these receipts as a coherent story:
what was asked, what was accepted, what inputs are permitted, what plan was reviewed,
what ran, what shipped, what drifted, how you will operate, and how you will review change.
