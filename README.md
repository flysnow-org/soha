# SOHA

SOHA is an enhanced library of Go(Golang) templates，it extends the capabilities of the Go template engine by custom function.

[简体中文](README_zh.md)

## Prerequisite

1. Go version 1.13+
2. Using Go Modules

## Quick start

```
# assume the following codes in example/hello.go  file
$ cat example/hello.go 
```

```go
package main

import (
	"github.com/flysnow-org/soha"
	"html/template"
	"log"
	"os"
)
func main() {

	sohaFuncMap := soha.CreateFuncMap()
	const templateText = `{{md5 .}}`

	tmpl, err := template.New("titleTest").Funcs(sohaFuncMap).Parse(templateText)
	if err != nil {
		log.Fatalf("parsing: %s", err)
	}

	err = tmpl.Execute(os.Stdout, 123456)
	if err != nil {
		log.Fatalf("execution: %s", err)
	}

}
```

```
➜  soha go run example/hello.go  
e10adc3949ba59abbe56e057f20f883e
```

## Golang Gin web framework example

```go
router := gin.Default()
    
sohaFuncMap := soha.CreateFuncMap()
router.SetFuncMap(sohaFuncMap)
```

## SOHA Functions List

#### int
Creates an int from the argument passed into the function.
```
int INPUT
```

#### string
Creates a string from the argument passed to the function
```
string INPUT
```

#### float
Creates a float from the argument passed into the function.
```
float INPUT
```

---

#### first
Slices an array to only the first N elements.
```
first LIMIT COLLECTION
```

#### last
slices an array to only the last Nth elements.
```
last INDEX COLLECTION
```

#### after
after slices an array to only the items after the Nth item.
```
after INDEX COLLECTION
```

#### complement
complement gives the elements of a collection that are not in any of the others.
```
COLLECTION | complement COLLECTION [COLLECTION]...
```

#### symdiff
symdiff returns the symmetric difference of two collections.
```
COLLECTION | symdiff COLLECTION
```

#### dict
Creates a dictionary from a list of key and value pairs.
```
dict KEY VALUE [KEY VALUE]...
```

#### echoParam
Prints a parameter if it is set.
```
echoParam DICTIONARY KEY
```

#### in
Checks if an element is in an array or slice–or a substring in a string—and returns a boolean.
```
in SET ITEM
```

#### intersect
Returns the common elements of two arrays or slices.
```
intersect SET1 SET2
```

#### union
Given two arrays or slices, returns a new array that contains the elements or objects that belong to either or both arrays/slices.
```
union SET1 SET2
```

#### isset
Returns true if the parameter is set.
```
isset COLLECTION INDEX

isset COLLECTION KEY
```

#### querify
Takes a set of key-value pairs and returns a query string to be appended to URLs.
```
querify KEY VALUE [KEY VALUE]...
```

#### shuffle
Returns a random permutation of a given array or slice.
```
shuffle COLLECTION
```

#### where
Filters an array to only the elements containing a matching value for a given field.
```
where COLLECTION KEY [OPERATOR] MATCH
```

#### append
append appends one or more values to a slice and returns the resulting slice.
```
COLLECTION | append VALUE [VALUE]...

COLLECTION | append COLLECTION
```

#### seq
Creates a sequence of integers.
```
seq LAST

seq FIRST LAST

seq FIRST INCREMENT LAST
```

#### uniq
Takes in a slice or array and returns a slice with subsequent duplicate elements removed.
```
uniq SET
```

---

#### md5
hashes the given input and returns its MD5 checksum.
```
md5 INPUT
```

#### sha1
hashes the given input and returns its SHA1 checksum.
```
sha1 INPUT
```

#### sha256
hashes the given input and returns its SHA256 checksum.
```
sha256 INPUT
```

---

#### base64
`base64Encode` and `base64Decode` let you easily decode content with a base64 encoding and vice versa through pipes.
```
base64Decode INPUT

base64Encode INPUT
```

#### jsonify
Encodes a given object to JSON.
```
jsonify INPUT
```

---

#### numFmt
Formats a number with a given precision using the requested negative, decimal, and grouping options. The options parameter is a string consisting of <negative> <decimal> <grouping>.
```
numFmt PRECISION NUMBER [OPTIONS [DELIMITER]]
```

#### add
Add adds two numbers
```
add INPUT1 INPUT2
```

#### sub
Subtracts two numbers.
```
sub INPUT1 INPUT2
```

#### mul
Multiplies two numbers.
```
mul INPUT1 INPUT2
```

#### div
Divides two numbers.
```
div INPUT1 INPUT2
```

#### mod
Modulus of two integers,returns `INPUT1 % INPUT2`.
```
div INPUT1 INPUT2
```

#### modBool
Boolean of modulus of two integers. Evaluates to `true` if `INT1%INT2` equals 0.
```
modBool INT1 INT2
```

#### math.Ceil
Returns the least integer value greater than or equal to the given number.
```
math.Ceil FLOAT
```

#### math.Floor
Returns the greatest integer value less than or equal to the given number.
```
math.Floor FLOAT
```

#### math.Round
Returns the nearest integer, rounding half away from zero.
```
math.Round FLOAT
```

#### math.Log
Log returns the natural logarithm of a number.
```
math.Log FLOAT
```

---

#### path.Base
Base returns the last element of a path.
```
path.Base PATH
```

#### path.Dir
Dir returns all but the last element of a path.
```
path.Dir PATH
```

#### path.Ext
Ext returns the file name extension of a path.
```
path.Ext PATH
```

#### path.Join
Join path elements into a single path.
```
path.Join ELEMENT...
```

#### path.Split
Split path immediately following the final slash.
```
path.Split PATH
```

---

#### reflect.IsMap
Reports if a value is a map.
```
reflect.IsMap INPUT
```

#### reflect.IsSlice
Reports if a value is a slice.
```
reflect.IsSlice INPUT
```

#### safeCSS
Declares the provided string as a known “safe” CSS string.
```
safeCSS INPUT
```

#### safeHTML
Declares a provided string as a “safe” HTML document to avoid escaping by Go templates.
```
safeHTML INPUT
```

#### safeHTMLAttr
Declares the provided string as a safe HTML attribute.
```
safeHTMLAttr INPUT
```

#### safeJS
Declares the provided string as a known safe JavaScript string.
```
safeJS INPUT
```

---

#### chomp
Removes any trailing newline characters.
```
chomp INPUT
```

#### findRE
Returns a list of strings that match the regular expression.
```
findRE PATTERN INPUT [LIMIT]
```

#### hasPrefix
Tests whether a string begins with prefix.
```
hasPrefix STRING PREFIX
```

#### lower
Converts all characters in the provided string to lowercase.
```
lower INPUT
```

#### upper
Converts all characters in a string to uppercase
```
upper INPUT
```

#### replace
Replaces all occurrences of the search string with the replacement string.
```
replace INPUT OLD NEW
```

#### replaceRE
Replaces all occurrences of a regular expression with the replacement pattern.
```
replaceRE PATTERN REPLACEMENT INPUT
```

#### split
splits a string into substrings separated by a delimiter
```
split STRING DELIM
```

#### substr
Extracts parts of a string from a specified character's position and returns the specified number of characters.
```
substr STRING START [LENGTH]
```

#### trim
Returns a slice of a passed string with all leading and trailing characters from cutset removed.
```
trim INPUT CUTSET
```

#### title
Converts all characters in the provided string to title case.
```
title INPUT
```

#### truncate
Truncates a text to a max length without cutting words or leaving unclosed HTML tags.
```
truncate SIZE INPUT
```

#### strings.HasSuffix
Determine whether or not a given string ends with the provided trailing suffix string.
```
strings.HasSuffix STRING SUFFIX
```

#### strings.Repeat
Returns a string consisting of count copies of the string s.
```
strings.Repeat INPUT COUNT
```

#### strings.RuneCount
Determines the number of runes in a string.
```
strings.RuneCount INPUT
```

#### strings.TrimLeft
Returns a slice of a given string with all leading characters contained in the cutset removed.
```
strings.TrimLeft CUTSET STRING
```

#### strings.TrimPrefix
Returns a given string s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.
```
strings.TrimPrefix PREFIX STRING
```

#### strings.TrimRight
Returns a slice of a given string with all trailing characters contained in the cutset removed.
```
strings.TrimRight CUTSET STRING
```

#### strings.TrimSuffix
Returns a given string s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.
```
strings.TrimSuffix SUFFIX STRING
```

---

#### dateFormat
Converts the textual representation of the datetime into the specified format.
```
dateFormat LAYOUT INPUT
```

#### now
Returns the current local time
```
now
```

#### time
Converts a timestamp string into a `time.Time` structure.
```
time INPUT
```

#### duration
Duration converts the given number to a time.Duration.
Unit is one of nanosecond/ns, microsecond/us/µs, millisecond/ms, second/s, minute/m or hour/h.
```
duration UNIT NUMBER
```

---

#### htmlEscape
Returns the given string with the reserved HTML codes escaped.
```
htmlEscape INPUT
```

#### htmlUnescape
Returns the given string with HTML escape codes un-escaped.
```
htmlUnescape INPUT
```

## Example

You can find a number of ready-to-run examples at [SOHA examples repository](example/)

## Acknowledgements

Thanks to [HUGO](https://github.com/gohugoio/hugo),SOHA extracted HUGO's template functions and Enhance them.

## License

SOHA is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.