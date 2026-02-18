# Names-only review paths

Reviewers should not browse the repo.

Start here:
- `fixtures/expected/case12_change_review_kit_smoke/kit_index.json`
- `fixtures/expected/case12_change_review_kit_smoke/manifest.sha256`

Then inspect only these kit inputs:
- `fixtures/input/case12_change_review_kit_smoke/kit/kit.json`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/change_request.json`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/acceptance.md`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/risk_rollback.md`
- `fixtures/input/case12_change_review_kit_smoke/kit/inputs/fixture_map.md`

Proof gate (must pass before shipping):
- `make verify`
