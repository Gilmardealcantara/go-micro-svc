# `Gilmardealcantara/github-actions/golang/test`

## Usage

```yaml
      - name: Unit Tests
        uses: Gilmardealcantara/github-actions/golang/test@main
        with:
          unit_tests_exclude: 'integration/tests' # skip integration tests
          coverage_threshold: 85

```

<!-- action-docs-inputs -->
## Description

Require run build before if there is some private repo

## Inputs

| name | description | required | default |
| --- | --- | --- | --- |
| `unit_tests_exclude` | <p>Pattern for exclude files folders using go list, ex: 'cmd|interfaces'</p> | `false` | `xxxxxxx` |
| `coverage_threshold` | <p>Test Coverage Threshold in %</p> | `false` | `80` |
| `coverage_file` | <p>Test Coverage txt file generated when you run tests</p> | `false` | `coverage.txt` |


