# gopt: The Go Path Tool

`gopt` is a tool for working with Go source files in a Go-sensitive way.
It is designed to make it easier for you to work with Go operations.

This is a tool to help with your...

- Shell scripting
- Makefiles
- TravisCI, CircleCI, and build pipeline config
- Daily life

## Usage
Run tests while ignoring `vendor/`

```
$ go test $(gopt pkgs)
```

Or for path-aware tools:

```
$ golint $(gopt paths)
```

### More

With a multi-path `$GOPATH`, find which path the current package is on:

```
⇒  gopt gopath
/Users/mbutcher/Code/Go
```

Get the name of the current package, even with a compound `$GOPATH`:

```
⇒  gopt name
github.com/Masterminds/glide
```

Get a list of subpackages in the project with `gopt pkgs`. Note that
this only returns packages that have source files (not including tests).

```
⇒  gopt pkgs
github.com/Masterminds/glide/action
github.com/Masterminds/glide/cfg
github.com/Masterminds/glide/dependency
github.com/Masterminds/glide/gb
github.com/Masterminds/glide/godep
github.com/Masterminds/glide/gpm
github.com/Masterminds/glide/importer
github.com/Masterminds/glide/msg
github.com/Masterminds/glide/path
github.com/Masterminds/glide/repo
github.com/Masterminds/glide/tree
github.com/Masterminds/glide/util
```

Or if you want to work with paths instead:

```
⇒  gopt paths
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/action
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/cfg
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/dependency
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/docs
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/gb
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/godep
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/gpm
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/importer
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/msg
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/path
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/repo
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/tree
/Users/mbutcher/Code/Go/src/github.com/Masterminds/glide/util
```

Note that this does not restrict to only paths with source!

## Installation

```
$ go get github.com/Masterminds/gopt
```

For development:

```
$ git clone THIS_REPO
$ glide install
```
