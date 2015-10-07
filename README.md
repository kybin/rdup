# rdup
Remove adjecent duplicated lines from stdin. Use it with pipe.

## usage
```bash
cat log.txt
[error] [00:00:00] hello
[error] [00:00:00] hello
[error] [00:00:00] hello
[error] [00:00:00] world
[error] [00:00:00] world
[error] [00:00:00] world
[error] [00:00:00] hello
[error] [00:00:00] world

# rdup will remove adjecent duplicated lines.
cat log.txt | rdup
[error] [00:00:00] hello
[error] [00:00:00] world
[error] [00:00:00] hello
[error] [00:00:00] world

# -n flag will let rdup remeber multiple lines.
cat log.txt | rdup -n=2
[error] [00:00:00] hello
[error] [00:00:00] world
```
