# pipeline-security-kit

[![Test rules](https://github.com/luan-evangelista/pipeline-security-kit/actions/workflows/test-rules.yml/badge.svg)](https://github.com/luan-evangelista/pipeline-security-kit/actions/workflows/test-rules.yml)

**Add a full security pipeline to any GitHub repository in ten lines of YAML.**

One reusable workflow that scans for leaked secrets, insecure code, vulnerable dependencies, risky configuration and vulnerable container images, plus a set of tested Semgrep rules for payment and fintech code that generic rulesets miss.

```yaml
jobs:
  security:
    uses: luan-evangelista/pipeline-security-kit/.github/workflows/security-scan.yml@<commit-sha>
    permissions:
      contents: read
      security-events: write
    with:
      codeql-languages: '["javascript-typescript"]'
      docker-context: .
```

That is the whole integration. Results show up in the repository's **Security** tab.

## What runs

| Job | Tool | Runs when | Fails the build when |
| --- | --- | --- | --- |
| Secrets | Gitleaks CLI, full history, checksum-verified binary | always | any secret is found |
| Code | Semgrep: registry rulesets + this kit's rules | always | an `ERROR`-severity rule matches |
| Code | CodeQL `security-extended` | `codeql-languages` is set | never (report only) |
| Dependencies and config | Trivy filesystem (CVEs + IaC/Dockerfile misconfig) | always | a finding in `fail-on` has a fix |
| Container | Docker build + Trivy image + Syft SBOM | `docker-context` is set | a finding in `fail-on` has a fix |

## Inputs

| Input | Default | What it does |
| --- | --- | --- |
| `semgrep-rulesets` | `p/default` | Registry rulesets added to the kit's rules, space-separated |
| `codeql-languages` | `[]` | JSON list of CodeQL languages; empty skips CodeQL |
| `docker-context` | `""` | Build context to build and scan; empty skips the container job |
| `dockerfile` | `Dockerfile` | Dockerfile path |
| `fail-on` | `CRITICAL,HIGH` | Severities that fail the dependency and image gates |
| `kit-ref` | `main` | Version of this kit's rules to use; pin it to a SHA |

## Custom rules

Generic rulesets are good at generic bugs. These rules target what hurts payment systems, and each one ships with unit tests (`tests/`) that run on every change.

| Rule | Languages | Catches |
| --- | --- | --- |
| `pan-in-logs` | TS/JS, Go | Card number, CVV or PIN passed to a logger (PCI DSS 3.3, 3.4) |
| `tls-verification-disabled` | TS/JS, Go | `rejectUnauthorized: false`, `NODE_TLS_REJECT_UNAUTHORIZED=0`, `InsecureSkipVerify` (PCI DSS 4.2.1) |
| `sql-string-building-go` | Go | SQL built with `fmt.Sprintf` or `+`, directly or through a variable |
| `weak-password-hash` | TS/JS, Go | MD5 and SHA-1 |
| `jwt-hardcoded-secret` | TS/JS, Go | JWTs signed or verified with a literal key |

```bash
pip install semgrep
semgrep --test --config rules/ tests/          # rule unit tests
semgrep scan --config rules/ path/to/your/code # use the rules locally
```

## Design choices

- **Gate only on what is actionable.** Builds fail on critical/high findings that have a fix. Everything else is visible but does not block, so developers do not learn to ignore the pipeline.
- **The pipeline is hardened too.** Actions pinned to commit SHAs, least-privilege `permissions`, `persist-credentials: false`, Harden-Runner on every job, downloaded binaries verified by SHA-256.
- **No license keys or paid accounts.** Everything runs on the default `GITHUB_TOKEN` with open-source tools.

## Rolling it out

Turning on blocking checks on day one in an existing codebase usually means a red pipeline nobody trusts. See [docs/ADOPTION.md](docs/ADOPTION.md) for a staged rollout and [docs/TRIAGE.md](docs/TRIAGE.md) for how to handle findings and exceptions.

## See it in action

- [secure-pipeline-demo](https://github.com/luan-evangelista/secure-pipeline-demo): the same layers applied to a TypeScript payments API.
- [vulnerable-to-secure](https://github.com/luan-evangelista/vulnerable-to-secure): a vulnerable Go API, what the pipeline catches, and the fixed version that passes.

## Want this set up and tuned for your team?

I do this as a fixed-scope service: pipeline setup, rule tuning for your stack and a triage session with your developers. **[levangelista.com.br](https://levangelista.com.br)**

## License

MIT
