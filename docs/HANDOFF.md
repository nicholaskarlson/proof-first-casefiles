# Handoff

A casefile is meant to be attachable to a client handoff or audit pack.

Each kit has a deterministic evidence bundle produced by this repo:
- `kit_index.json` (normalized inventory + hashes)
- `manifest.sha256` (sha256 lines over the kit file tree)

If a kit violates the contract (expected-fail), the output is:
- `error.txt` only
