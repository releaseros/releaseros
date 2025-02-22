# Releaseros Documentation

This is the documentation of [Releaseros](https://github.com/releaseros/releaseros).

## Configuration reference

The configuration file is a [YAML](https://yaml.org/) file.

The default path for the config file is `.releaseros.yaml` that is placed in the working directory.
All config file names supported by order of priority are:

* `.releaseros.yaml`
* `.releaseros.yml`
* `releaseros.yaml`
* `releaseros.yml`

As soon as one of these files is found, the config file is loaded.

The `--config` option allows you to specify a custom configuration file.

### Parameters

#### `initial_release_message`

The message that will be used if no previous tag exist.

Example: `Initial Release`

#### `sort`

The sorting order for the commits.
Either `asc` (ascending) or `desc` (descending).

#### `filters`

The filters to apply to the commits.

* `exclude`: a list of regular expressions. If a commit matches one of these regular expressions,
it will be excluded from the release note.

Example:

```yaml
filters:
  exclude:
    - '^test:'
    - '^ci:'
    - 'merge conflict'
    - Merge pull request
    - Merge remote-tracking branch
    - Merge branch
```

#### `categories`

The categories to use for the commits.

* `title`: the title of the category.
* `regexp`: a regular expression that will be used to match the commits.
* `weight`: An integer representing the weight of the category. The heavier the weight, the more the category will be at the bottom of the release note.

Example:

```yaml
categories:
  - title: 'Features'
    regexp: '^.*?feat(\([[:word:]-]+\))??!?:.+$'
    weight: 10
  - title: 'Fixes'
    regexp: '^.*?fix(\([[:word:]-]+\))??!?:.+$'
    weight: 20
  - title: 'Documentation'
    regexp: '^.*?docs(\([[:word:]-]+\))??!?:.+$'
    weight: 30
  - title: Others
    weight: 9999
```

#### `footer`

The footer of the release note.

Two placeholders are available:

* `PreviousTag`: the previous tag.
* `LatestTag`: the latest tag.

Functions may be used:

* `{{ env "FOO" }}`: the value of the environment variable `FOO`. If the environment variable is not set, an empty string will be used.
* `{{ envOr "FOO" "fallback value" }}`: the value of the environment variable `FOO` or `fallback value` if `FOO` is not set.
* `{{ envOrError "FOO" }}`: the value of the environment variable `FOO` or an error if `FOO` is not set.

Example:

```
**Full Changelog**: https://{{ env "HOST" }}/my-project/-/compare/{{ .PreviousTag }}...{{ .LatestTag }}
```
