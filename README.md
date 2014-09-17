# sprout

An _experimental_ command line interface for [Pivotal Sprout](https://github.com/pivotal-sprout)

Currently a thin wrapper around [soloist](https://github.com/mkocher/soloist)

## Installation

Downloads and installation instructions can be found in the notes for the [latest release](https://github.com/hiremaga/sprout/releases/latest).

## Roadmap

### Parity with today's workflow

- [x] Ensure system Ruby and shell out to soloist
- [x] Automatically install soloist if it isn't already installed
- [x] Automatically install librarian-chef if it isn't already installed
- [ ] Run `soloist` via  sprout `sprout soloist ...`
- [ ] Run `librarian-chef` from sprout instead of soloist `sprout libarian-chef ...`
- [ ] Use dedicated gem home `~/.soloist/gems`

### Future improvements & deas

- [ ] ~~Use chef-omnibus instead of system ruby?~~ no need, omnibus is BIG and system ruby is sufficient
- [ ] Self-update
- [ ] Generate sprout-wrap
- [ ] Namespace soloist commands under `sprout solo ...` so sprout's interface can vary from soloist's
- [ ] homebrew tap/curl for easy installation
