# dashphrase

A command line tool for generating memorable passphrases.

## What is dashphrase?

A *Dashphrase* is a passphrase composed of a number of dictionary words, separated by dashes, followed by a number of random digits.

## How to use this tool?

Run it in the terminal:

```bash
Usage:

  -v        Print version information and exit
  -h        Print usage and exit

  -d int    Number of digits to append (default 3)
  -w int    Number of words in the phrase (default 3)

  -c        Show character count
  -n        Use a name in the phrase
  -u        Use underscores instead of dashes
```

By default, this tool will generate a passphrase with 3 words, separated by dashes, followed by 3 random digits. You can customize the number of words and digits using the `-w` and `-d` flags respectively.

A passphrase generated this way should be of sufficient length and complexity to use as an easy to remember password in most systems.

Most default dashprhases will default to:

- 20+ characters
- At least one upper-case character
- 3+ numbers
- 3+ special characters (dashes)

You can switch the dashes to underscores using the `-u` flag if your system does not consider dashes to be special characters.

Output Examples:

```bash
Excursion-tackiness-hypnotism-701
Showplace-audition-aviation-745
Glorify-squishy-wrist-346
Praising-straggler-lazy-296
Overbill-casing-trend-116
Foe-revenue-upper-610
Surrender-jackpot-showman-297
Generic-letter-aground-470
Cufflink-sublevel-fiddle-249
Magenta-dropbox-enunciate-785
```

## Installation

### Milti-Platform:

You can install dashphrase using `go install`:

```bashbash
go install github.com/maciakl/dashphrase@latest
```

### macOS and Linux:

Use [grab](https://github.com/maciakl/grab):

```bash
grab maciakl/dashphrase
```

### Windows:

Use [scoop](https://scoop.sh):

```bash
scoop add maciak https://github.com/maciakl/bucket
scoop update
scoop install dashphrase
```




