# chango
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/BodeSpezial/chango/trunk?style=flat-square)
![GitHub release (latest SemVer including pre-releases)](https://img.shields.io/github/v/release/BodeSpezial/chango?include_prereleases&sort=semver&style=flat-square)
![GitHub Release Date](https://img.shields.io/github/release-date-pre/BodeSpezial/chango?style=flat-square)

A command-line tool to automatically add and generate changelogs without the headache of permanent merge-conflicts

---

While working on a project a collegue and I recently discovered that it could be a real pain to keep a proper changelog if multiple people are often  merging onto one branch. So we decided to use the same approach as described in this article. We're using yarn and a ts script currently but to implement the same approach easily in future projects I'll try to create a cli-tool for this in go.

## Installation

As of now it's only in a development state and as so you have to clone and it install it manually

## Usage

### View
`chango view` prints out the changelog that will be created if you run `chango generate`. It's basically a dry run.

# TODO
- remove need for extra enter when opening vim
- Append entry to the right section if a changelog for the current branch already exists
