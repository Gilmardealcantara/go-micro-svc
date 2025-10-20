# `Gilmardealcantara/github-actions/golang/coverage`

## Usage

```yaml
    - name: Check Coverage
      uses: Gilmardealcantara/github-actions/golang/covera@main
      with:
        coverage_threshold: 85
        coverage_file: './tests/coverage.txt'
```

<!-- action-docs-inputs -->
## Description

Validate code coverage
- Require run test before to generate generate coverage.txt file

## Inputs

| name | description | required | default |
| --- | --- | --- | --- |
| `coverage_threshold` | <p>Test Coverage Threshold in %</p> | `false` | `80` |
| `coverage_file` | <p>Test Coverage txt file generated when you run tests</p> | `false` | `coverage.txt` |

