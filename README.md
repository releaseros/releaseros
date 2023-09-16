# Releaseros - Release Note Generator

[![Tests](https://github.com/releaseros/releaseros/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/releaseros/releaseros/actions?workflow=tests)
[![License](https://img.shields.io/github/license/releaseros/releaseros)](/LICENSE)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/releaseros/releaseros)](https://github.com/releaseros/releaseros/releases/latest)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-blue.svg?style=flat)](https://conventionalcommits.org)

🚧 WIP 🚧

Generate a release note based on Git repository.

If you need to generate a release note in a flexible way without too much work this may be the tool for you.

## Quick Start

### Basic

Initialize an example config file named `.releaseros.yaml`:

```bash
$ releaseros init
```

Generate the release note:

```bash
$ releaseros generate
```

That's it!

## Prerequisites

You need to have [Git](https://git-scm.com/) installed.

## Installation

### Binary

You will find the binary for your platform in [the latest release page](https://github.com/releaseros/releaseros/releases/latest).

> If you think that a platform is missing feel free to propose it!

### Docker

You may use the latest tag from (you should pin a specific version):
* [Dockerhub](https://hub.docker.com/r/releaseros/releaseros): `releaseros/releaseros:latest`
* [Github Container Registry](https://github.com/releaseros/releaseros/pkgs/container/releaseros): `ghcr.io/releaseros/releaseros:latest`

## Examples

### Gitlab integration

Let's assume you have the following configuration `.releaseros.yaml`:

```yaml
initial_release_message: |
  Initial Release
sort: asc
filters:
  exclude:
    - '^test:'
    - '^ci:'
    - 'merge conflict'
    - Merge pull request
    - Merge remote-tracking branch
    - Merge branch
categories:
  - title: 'Features'
    regexp: '^.*?feat(\([\w'-]+\))??!?:.+$'
    weight: 10
  - title: 'Fixes'
    regexp: '^.*?fix(\([\w'-]+\))??!?:.+$'
    weight: 20
  - title: 'Documentation'
    regexp: '^.*?docs(\([\w'-]+\))??!?:.+$'
    weight: 30
  - title: Others
    weight: 9999
footer: |
  **Full Changelog**: https://gitlab.com/myrepo/-/compare/{{ .PreviousTag }}...{{ .LatestTag }}
```

You could combine the [release feature of Gitlab](https://docs.gitlab.com/ee/ci/yaml/#release) with [releaseros](https://github.com/releaseros/releaseros) in a `.gitlab-ci.yml` like:

```yaml
stages:
  - generate
  - release

generate:
  stage: generate
    image:
    name: releaseros/releaseros:latest # remember to always pin a specific version
    entrypoint: ['']
  script:
    - releaseros generate > releasenote.md
  rules:
    - if: $CI_COMMIT_TAG
  variables:
    # Disable shallow cloning so that releaseros can diff between tags.
    GIT_DEPTH: 0
  artifacts:
    paths:
      - releasenote.md

release_job:
  stage: release
  image: registry.gitlab.com/gitlab-org/release-cli:v0.15.0
  rules:
    - if: $CI_COMMIT_TAG
  dependencies:
    - generate
  script:
    - echo "Running the release job."
  release:
    tag_name: $CI_COMMIT_TAG
    name: '$CI_COMMIT_TAG'
    description: './releasenote.md'
```

And you will obtain something like the following image

![releaseros combined with gitlab](/.github/img/releaseros-gitlab-release-example.png)

## TODOs

- [ ] logo of Gene the Releaseros
- [ ] write documentation

## Contributing

See the [contributing guide](CONTRIBUTING.md).

## Inspirations

During a long time I used a minimal [awk](https://www.gnu.org/software/gawk/manual/gawk.html) script.

Creating a tool that match my expectations was on my mind from a long time.

Then one day, I discovered [GoReleaser](https://github.com/goreleaser/goreleaser).
I was blown away by it!

On other projects that cannot used it (not Go projects), I was frustated by either having to manually create the release note, or using my good old `awk` script.

So for fun and to continue my learning of [Go](https://go.dev/) this project was created!
It is based on what my good old `awk` script did and on what [GoReleaser](https://github.com/goreleaser/goreleaser) have to propose when generating a release note.
