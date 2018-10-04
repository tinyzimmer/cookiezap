## CookieZap

Cookiezap deletes all cookies for the specified browser that match the search terms provided on the command line.
Sometimes you may need to restart your browser for the changes to take effect.

## Usage

```bash
$> ./cookiezap

Usage of ./cookiezap:

  -c    Clean Chrome cookies
  -f    Clean Firefox cookies

./cookiezap [args] <search_terms...>

Example:
$> cookiezap -f reddit
10 cookie(s) related to reddit deleted from Mozilla Firefox
```

## Build from source

### Build Requirements
 - go >= 1.10.x

```bash
$> go get github.com/tinyzimmer/cookiezap
$> go install github.com/tinyzimmer/cookiezap
```
