# Line Length Checker

A simple pre-commit hook to check for lines exceeding a specified character length in text files.

## Installation

Add this to your `.pre-commit-config.yaml`:

```yaml
- repo: https://github.com/scottnuma/check-line-length
  rev: v0.1.0
  hooks:
  - id: check-line-length
    args: ["--max-line-length=75"] # Optional, default is 80
    files: some/directory/.
```

