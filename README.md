[![⚗️Check](https://github.com/kijimaD/act/actions/workflows/check.yml/badge.svg)](https://github.com/kijimaD/act/actions/workflows/check.yml)

# act

<img src="https://user-images.githubusercontent.com/11595790/193450591-6b681517-3b5a-4dd4-ac04-5dce9b209882.png" width="40%" align=right>

act is curating your github activity tool.

# Install

```sh
$ go install github.com/kijimaD/act@main
```

# How to use

set `.act.yml`

```yml
outType: file
outFile: ./README.md
```

and prepare GitHub API token, run

```shell
$ GH_TOKEN=your_token... act
```

# Docker run

```shell
docker run --rm
           -it
           -v "${PWD}":/workdir \
           ghcr.io/kijimad/act:latest
```