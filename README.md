# headtail

Commandline utility to read and display head and tail of a file.

Combination of `head` and `tail`.

## Usage

```
$ perl -le 'print for 1..1100000' | headtail -n 4
1
2
3
4
1099997
1099998
1099999
1100000
```

## Rationale

Occasionally, it is useful to look at the extremes of sorted data.
