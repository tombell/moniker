# moniker

A command-line tool to rename MP3 files according to ID3v2 tags.

## Installation

To get the most up to date binaries, check [the releases][releases] for the
pre-built binary for your system.

[releases]: https://github.com/tombell/moniker/releases

You can also `go get` to install from source.

    go get -u github.com/tombell/moniker/cmd/moniker

## Usage

Using `moniker` is very simple, and only has 1 optional flag for specifying the
filename formatting. By default the format will be `{artist} - {title}`.

If you wish to use another format, you can use the `--format` flag, and the
formatter options.

Currently supported options:

  - `{album}`
  - `{artist}`
  - `{genre}`
  - `{title}`
  - `{year}`

So to rename directory of MP3 files, including the artist, title, and genre the
command would be as follows.

    moniker --format "{genre} - {artist} - {title}" ~/Music/New

That's all!
