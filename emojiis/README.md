[![Go Reference](https://pkg.go.dev/badge/github.com/vladimirvivien/emojiis.svg)](https://pkg.go.dev/github.com/vladimirvivien/emojiis)

# emojiis

Emojiis is a Go module that exposes a programmatic API to search and retrieve emoji icons using descriptive terms.

## Installation
To get started with the module, use the `go get` command to pull down the packages:

```
go get github.com/vladimirvivien/emojiis@latest
```

## Usage
Here is a simple example of how to use the API:

```go
emojis := search.ByDescription(search.Params{
    Include: []string{"face"}, 
    Exclude: []string{"smile", "laugh", "grin", "upside-down"}
})
```

Returns:

```
[]string{"🐵", "🐶", "🐱", "🐯", "🦊"}
```
