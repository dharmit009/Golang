# Strings

Strings are sequences of bytes.

ASCII - American Standard Code for Information Exchange.

Character coding - each character is associated with an (7) 8-bit
number.

'A' = 0x41

Unicode is a character code of 32-bits.

Now the default in Go is UTF-8 character.

*What are Runes?*

> Runes are just code points in Go. i.e for A is 41.
> Each byte of a string is a Rune.

## String packages:

* Runes are divided into many different categories.
* Provides a set of functions to test categories of runes.

**For Example:**

IsDigit( r rune )
IsSpace( r rune )
IsLower( r rune )
IsPunct( r rune )
IsLetter( r rune )

* Some functions perform conversion:

1. ToUpper( r rune );
1. ToLower( r rune );

* String Search Functions

1. Compare ( a,b )
1. Contains( s, substr )
1. HasPrefix( s, substr )
1. Index ( s, substr )

Strings are immutable, but modified strings can be returned.

* Replace ( s, old, new, n):

Replace returns a copy of the strings with the first `n` instances of
old replaced by new.

* TrimSpace(s):

Returns a new string with all leading and trailling spaces removed.

* Atoi( s ):

Converts string s to int

* Itoa( s ):

Converts int ( base 10 ) to string.

* FormatFloat( f, fmt, prec, bitSize ):

Converts floating point number to a string.

* ParseFloat ( s, bitSize ):

Converts a string to floating point number.

















