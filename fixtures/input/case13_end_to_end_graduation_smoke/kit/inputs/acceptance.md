# Acceptance criteria (graduation kit)

Accepted when a reviewer can validate the end-to-end proof chain **by opening only named artifacts**:

- Deterministic outputs exist: `kit_index.json` and `manifest.sha256`.
- The kit contains a coherent receipt chain: contract → strict decode → plan → execution bundle → drift → deploy evidence → ops guardrails → change review.
- No secrets, no timestamps, no generated randomness.
- This is a **green path** smoke fixture; expected-fail posture lives in sibling fixtures.

A reviewer should be able to say:
“I can see what you planned, what you ran, what you shipped, what drifted, and how you will operate and review change — all with deterministic receipts.”
