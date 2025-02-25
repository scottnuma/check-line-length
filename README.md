# Line Length Checker

A pre-commit hook to check for lines exceeding a specified character length in text files.

## Installation

Add this to your `.pre-commit-config.yaml`:

```yaml
-   repo: https://github.com/scottnuma/check-line-length
    rev: v0.1.0
    hooks:
    -   id: line-length-checker
        args: [--max-line-length=80]  # Change as needed
```
