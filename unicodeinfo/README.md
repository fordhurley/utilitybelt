unicodeinfo
============

Prints out unicode information for each character on stdin, or from the
provided arguments.

```
Usage: unicodeinfo [-h | "String to decode!"]
	-h  print this help
```


Install
--------

```
go get -u github.com/fordhurley/utilitybelt/unicodeinfo
```


Examples
---------

```
unicodeinfo 👩‍👩‍👧‍👦
0: U+1F469 '👩' (WOMAN)
4: U+200D (ZERO WIDTH JOINER)
7: U+1F469 '👩' (WOMAN)
11: U+200D (ZERO WIDTH JOINER)
14: U+1F467 '👧' (GIRL)
18: U+200D (ZERO WIDTH JOINER)
21: U+1F466 '👦' (BOY)
```

```
$ $ unicodeinfo < README.md
0: U+0075 'u' (LATIN SMALL LETTER U)
1: U+006E 'n' (LATIN SMALL LETTER N)
2: U+0069 'i' (LATIN SMALL LETTER I)
3: U+0063 'c' (LATIN SMALL LETTER C)
4: U+006F 'o' (LATIN SMALL LETTER O)
5: U+0064 'd' (LATIN SMALL LETTER D)
6: U+0065 'e' (LATIN SMALL LETTER E)
7: U+0069 'i' (LATIN SMALL LETTER I)
8: U+006E 'n' (LATIN SMALL LETTER N)
9: U+0066 'f' (LATIN SMALL LETTER F)
10: U+006F 'o' (LATIN SMALL LETTER O)
...
```
