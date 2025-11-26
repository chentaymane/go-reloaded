🌟 go-reloaded
A Powerful Text Formatter & Auto-Correct Tool Written in Go

🚀 Overview

go-reloaded is a command-line tool that reads a text file, applies a series of smart transformations, and outputs a fully corrected, formatted and human-friendly version.

It was created for the Zone01 Oujda curriculum and demonstrates advanced use of:
✔ string manipulation
✔ pattern detection
✔ Go file system (os)
✔ algorithmic thinking
✔ custom text processing logic

✨ Features
🔢 Number Conversions
Tag	Description
(hex)	Converts previous hexadecimal word → decimal
(bin)	Converts previous binary word → decimal

Example

1E (hex) → 30
10 (bin) → 2

🔠 Case Transformations
Tag	Action
(up)	UPPERCASE
(low)	lowercase
(cap)	Capitalize first letter

Supports repetitions:

(up, 3) → transforms previous 3 words into UPPERCASE

✏️ Punctuation Normalization

Automatically fixes bad punctuation spacing:

❌ Hello ,world !!
✔ Hello, world!!

Supports multi-punctuation:

... !! !? ,, ...

📝 Smart Quote Handling (' ')

Quotes are always placed exactly around the intended word(s):

' awesome ' → 'awesome'


Multiple-word quotes also work:

' I am the best ' → 'I am the best'

🅰️ → 🅰️🅽 Article Correction

Automatically turns “a” into “an” when followed by:
a e i o u h

Example:

a amazing → an amazing
a honest → an honest

📦 Installation
git clone https://github.com/<yourusername>/go-reloaded.git
cd go-reloaded

▶️ Usage
Run the program:
go run . input.txt output.txt

📘 Examples
✔ Case + Punctuation

Input:

it (cap) was the best of times ,and the worst of times (up) !!


Output:

It was the best of times, and the worst of TIMES!!

✔ Hex + Binary

Input:

Simply add 42 (hex) and 10 (bin).


Output:

Simply add 66 and 2.

✔ Quote Fixing

Input:

I am ' awesome ' and ' very cool ' today.


Output:

I am 'awesome' and 'very cool' today.

✔ “a” → “an”

Input:

This is a amazing project.


Output:

This is an amazing project.

🧪 Unit Tests

You are encouraged to add tests for:

Hex/Bin conversions

Text case transformations

Quotation logic

Punctuation rules

"a/an" handling

📚 What You Learn

Mastering Go slices & strings

Custom lexing/tokenizing

OS file reading/writing

Edge-case handling

Building a real tool from scratch

🤝 Contributing

Feel free to:

Open an Issue

Submit a Pull Request

Suggest improvements ❤️

🎯 Author

Project completed by Aymane Chent
Part of Zone01 Oujda — Go Reloaded Project.
