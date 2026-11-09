# Triaging findings

Every finding ends in one of four states. "Ignored" is not one of them.

| State | When | What to do |
| --- | --- | --- |
| **Fixed** | Real issue with a fix | Fix it in the same PR when possible |
| **False positive** | The tool is wrong for this code | Inline suppression with a reason |
| **Accepted risk** | Real, but no fix exists yet or the impact is contained | Documented exception with an owner and an expiry date |
| **Not applicable** | Test fixtures, generated code, vendored code | Exclude the path in the tool config, with a comment |

## Suppressing a false positive

Always say why. A suppression without a reason will be questioned in review, and should be.

```ts
const fixture = "tok_4f9a2c7e1b3d5f80"; // gitleaks:allow (fake token used only in tests)
```

```go
db.Query(query) // nosemgrep: sql-string-building-go (query is a compile-time constant)
```

## Fix deadlines

A common starting point, to be agreed with the team:

| Severity | Deadline |
| --- | --- |
| Critical | 7 days |
| High | 30 days |
| Medium | 90 days |
| Low | Next planned maintenance |

Teams under PCI DSS must install critical security patches within one month of release (requirement 6.3.3) and set the other deadlines through their own risk analysis.

## Exceptions

Track each exception where the team already tracks work (an issue with a `security-exception` label works) with:

- the finding and why it cannot be fixed now
- the compensating control, if any
- an owner
- an expiry date, after which the finding is blocking again
