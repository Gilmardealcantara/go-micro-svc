# `Gilmardealcantara/github-actions/golang/build`

## Usage

```yaml
      - name: Application Build
        uses: Gilmardealcantara/github-actions/golang/build@main
        with:
          user: ${{ secrets.GH_PACKAGE_USERNAME }}
          token: ${{ secrets.GH_PACKAGE_TOKEN }}
```

<!-- action-docs-inputs -->
## Description

Runs Go within internal folder

## Inputs

| name | description | required | default |
| --- | --- | --- | --- |
| `token` | <p>Git package token</p> | `true` | `""` |
| `user` | <p>Git package username</p> | `true` | `""` |
| `main_path` | <p>Path of your main file</p> | `true` | `.` |
| `output_bin_path` | <p>If provided, the app bin wil base save in this path</p> | `false` | `./bin` |
| `output_bin_name` | <p>If provide, the upload-artifact will receive this name </p> | `false` | `""` |

