# ascii-art

Small Go utility that renders input text as ASCII-art using a fixed font defined in `standard.txt`.

## Overview
- Reads a font file (`standard.txt`) with glyphs for printable ASCII characters (starting at ASCII 32).
- Each glyph block is 9 lines; renderer uses rows 0–7 to print an 8-line ASCII-art output per character.
- Implementation: `ascii-art/main.go`.

## Requirements
- Go 1.22 (see `go.mod`).

## Usage
Build:
```sh
go build -o ascii-art
```

Run (use literal `\n` for line breaks in the single argument):
```sh
# with go run
go run main.go "HELLO\\nWORLD"

# or with built binary
./ascii-art "Hello, World!\\nThis is ascii-art"
```

Notes:
- `standard.txt` must be in the same directory as the binary or when running `go run`.
- Input is a single command-line argument; use `\n` (backslash + n) to split lines.
- Supported characters: printable ASCII starting at space (ASCII 32). Characters outside this range may produce incorrect output or errors.

## Troubleshooting
- If index panics or wrong output occur, ensure `standard.txt` is intact and contains the expected number of lines for all glyphs.