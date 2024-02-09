[![Go Reference](https://pkg.go.dev/badge/github.com/vladimirvivien/emojiis.svg)](https://pkg.go.dev/github.com/vladimirvivien/emojiis)

# emojiis

Emojiis is a Go module that exposes a programmatic API to search and retrieve emoji icons using descriptive terms.

## Version 2.0.0 released

Version 2.0 uses a JSON file (from https://github.com/milesj/emojibase) 
for its dataset, containing all commonly used emojis.

The API has been updated (if you use v1.xx, some functions may break after you update).

Features:
* Ability to search the full set of common emojis
* Support for inclusion and exclusion search terms
* Ability to search by emoji tags

 See our API [documentation](https://pkg.go.dev/github.com/vladimirvivien/emojiis)

## Installation
To get started with the module, use the `go get` command to pull down the packages:

```
go get github.com/vladimirvivien/emojiis/v2@v2.0.0
```

## Usage
Here is a simple example of how to use the API:

```go
emojis := search.All(search.Params{
    Include: []string{"face"}, 
    Exclude: []string{"smile", "laugh", "grin", "upside-down"}
})
```

Returns:
```
[]string{"🐵", "🐶", "🐱", "🐯", "🦊"}
```

Or, search by tags only:

```go
emojis := search.ByTags("fruits")
```

Returns:
```go
[]string{"🍇", "🍈", "🍉", "🍊", "🍋"}
```


