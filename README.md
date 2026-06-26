# ASCII Art

A simple and efficient command-line application written in Go that converts plain text into large ASCII art banners using a bitmap font file.

Each printable ASCII character is represented by an 8-line pattern stored in a font definition file (`thinkertoy.txt`). The program reads user input, maps each character to its corresponding glyph, and renders the complete banner directly in the terminal.

---

## Features

- Convert plain text into ASCII art
- Supports all printable ASCII characters
- Handles multi-line input using `\n`
- Lightweight and dependency-free
- Fast terminal rendering
- Simple CLI interface

---

## Project Structure

```
.
├── main.go
├── thinkertoy.txt
└── README.md
```

---

## Requirements

- Go 1.20 or newer
- `thinkertoy.txt` font file located in the project root

---

## Installation

Clone the repository:

```bash
git clone https://github.com/boanerges-dev/ascii-art.git
```

Navigate into the project directory:

```bash
cd ascii-art
```

---

## Usage

```bash
go run . "Hello"
```

or

```bash
go run main.go "Hello"
```

---

## Multi-line Input

Use `\n` to create multiple lines.

Example:

```bash
go run . "Hello\nWorld"
```

Output:

```
<ASCII representation of Hello>

<ASCII representation of World>
```

---

## How It Works

The program loads the bitmap font file:

```
thinkertoy.txt
```

Each printable ASCII character occupies **9 lines** in the font file:

- 8 lines for the glyph
- 1 empty separator line

For every character in the input:

1. Convert the character to its ASCII value.
2. Normalize the value by subtracting 32 (space is the first printable character).
3. Compute its starting position inside the font file.
4. Print the corresponding 8 rows.
5. Repeat for every character.

The lookup index is calculated as:

```go
index := (int(char)-32)*9 + row
```

---

## Example

Input

```bash
go run . "Go"
```

Output

```
 ██████      ██████
██          ██    ██
██  ███     ██    ██
██    ██    ██    ██
 ██████      ██████
```

*(Output depends on the contents of `thinkertoy.txt`.)*

---

## Error Handling

The application validates several error conditions:

- Missing command-line argument
- Empty input
- Missing font file
- File read errors

Example:

```text
Usage: go run main.go <input>
```

or

```text
Error reading file: open thinkertoy.txt: no such file or directory
```

---

## Limitations

- Supports only printable ASCII characters (32–126)
- Requires `thinkertoy.txt` to be present
- Does not validate unsupported Unicode characters

---

## Implementation Details

The application uses only Go's standard library:

- `fmt`
- `os`
- `strings`

No third-party dependencies are required.

---

## Future Improvements

- Multiple banner styles
- Colored terminal output
- Alignment options (left, center, right, justify)
- Output to file
- Unicode support
- Custom font selection
- Flag-based CLI (`-output`, `-font`, `-color`)

---

## Example Font Layout

Each character occupies nine consecutive lines:

```
Character 32 (Space)
Line 1
Line 2
...
Line 8
(blank)

Character 33 (!)
...
```

---

## Performance

Time Complexity:

```
O(n × 8)
```

where **n** is the number of characters in the input.

Space Complexity:

```
O(f)
```

where **f** is the size of the loaded font file.

---

## Contributing

Contributions are welcome.

To contribute:

1. Fork the repository.
2. Create a feature branch.

```bash
git checkout -b feature/my-feature
```

3. Commit your changes.

```bash
git commit -m "Add new feature"
```

4. Push to your fork.

```bash
git push origin feature/my-feature
```

5. Open a Pull Request.

---

## License

This project is released under the MIT License.

Feel free to use, modify, and distribute it for personal or commercial projects.

---

## Author

Developed in Go as a command-line ASCII art renderer for educational purposes.