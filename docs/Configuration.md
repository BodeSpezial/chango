# Configuration

### Editor (string)

The editor to use to edit the entries.

Default: vim
Supported: vim, nano

### Path to changelogs (string)

The path under which your changelog files are stored. This shouldn't be something you have to change under normal circumstances.

Default: changelog

## Git

This section holds the settings for your Version Control System (VCS).

### autocommit (bool)

Wether to commit your entry automatically or not. I added this because I kept typing
```shell
git add changelog/
git commit -m "Added changelog"
git push
```
after adding a new entry.
<!-- TODO: Add this flag
If you want to generally use this feature but disable it temporarely (bc you are still working on those changes and want to have the changelog in one commit) you can use `--auto-commit false` or the shorthand `-nac` (abbrev. for no automatic committing). -->

Default: true
Supported: true, false

### repo_url (string)

Default: empty

## Changelog

In this section you can personalize the generated changelog file.

### filename (string)

Default: Changelog

### filetype (string)

This will also determine the style in which your changelog is generated. So be aware, that only supported formats can be outputted. If the used format isn't supported `md` will be used.

Default: md
Supported: md, markdown

