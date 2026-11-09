# Rolling out the security pipeline

The goal is a gate developers trust. That takes a few weeks of tuning, not a single pull request.

## Week 1: see everything, block nothing

- Add the workflow with `fail-on: CRITICAL` and let it run on `main` and on pull requests.
- Do **not** make the checks required yet.
- Collect the baseline: how many findings per layer, and which are real.

## Week 2: clean up and tune

- Rotate and remove any real secret Gitleaks found. Deleting the line is not enough: it is still in git history and must be considered leaked.
- Upgrade dependencies with critical/high CVEs that have a fix.
- Mark false positives with an inline suppression **and a reason** (see [TRIAGE.md](TRIAGE.md)).
- Turn off Semgrep rules that are noisy for this codebase, and write down why.

## Week 3: start gating

- Make the `Security` checks **required** in branch protection.
- Move `fail-on` to `CRITICAL,HIGH`.
- Agree on fix deadlines per severity with the team lead.

## Ongoing

- The weekly scheduled run catches CVEs published for code that did not change.
- Dependabot keeps actions and dependencies current, including the pinned SHAs.
- Review open exceptions every month; expired exceptions become blocking again.
