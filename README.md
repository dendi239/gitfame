# gitfame

Command line tool for analyzing git-repo

## Install

For simple installation you need only
```shell
$ go install github.com/dendi239/gitfame@latest
```

You can build from source using following:
```shell
$ git clone https://github.com/dendi239/gitfame.git
$ cd gitfame
$ go install ./...
```

### Troubleshooting

Both of methods described above **require go 1.16** and will install executable to `$GOPATH/bin`, or `$GOBIN` so make sure it's added to `$PATH` (following example assumes that you use `zsh`)
```shell
$ echo '# Go lang
export GOPATH="$HOME/go"
export PATH="$GOPATH/bin:$PATH"
' >> ~/.zshrc
```

## Usage

```text
gitfame calculates (lines, files, commits) for selected revision

Usage:
  gitfame [flags]

Flags:
      --exclude strings       glob patterns to exclude from counting
      --extensions strings    extensions to count, e.g. '.go,.md'
      --format string         format, one of: "tabular", "json", "csv", "json", "json-lines" (default "tabular")
  -h, --help                  help for gitfame
      --languages strings     languages to count, e.g. 'go,markdown'
      --order-by string       key for result sorting, one of: "lines", "commits", "files" (default "lines")
  -p, --progress              shows actual progress during analyzing
      --repository string     path to repository (default ".")
      --restrict-to strings   if specified, every counted file should match at least one of these globs
      --revision string       revision to analyze (default "HEAD")
      --use-committer         Using committer instead of author
```
