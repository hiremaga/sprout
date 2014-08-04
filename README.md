# sprout

An _experimental_ command line interface for [Pivotal Sprout](https://github.com/pivotal-sprout)

Currently a thin wrapper around [soloist](https://github.com/mkocher/soloist)

Ideas:

- [x] Ensure system ruby and shell out to soloist
- [ ] Automatically install soloist if it isn't already installed
- [ ] Automatically install librarian-chef if it isn't already installed
- [ ] Run librarian-chef from sprout instead of soloist
- [ ] Use chef-omnibus instead of system ruby?
- [ ] Self-update
- [ ] Generate sprout-wrap
- [ ] Namespace soloist commands under `sprout solo ...` so sprout's interface can vary from soloist's
