# Scripting
## Info
The hyperscript (hs) interpreter is a separate executable, which is only imported by hypterm. It can also be used as a stand-alone program from within other shells. 

## Example
```js
// Comments are started with two slashes
// Generally speaking, the rule is: one command per line (Exceptions: piping)
run firefox

// Commands can be chained and piped, however:
ls |> filter "*.pdf"

// Variables always have a leading $-sign:
$x = 4

// These are all the types of hyperscript
1
1.0
"Strings"
true, false
[1, false, "true"]
{1: "a", "b": 2}

// If conditions are always comparisons
if $x {

} elseif $x == 2 {

} else {

}

// For-loops are declared as follows:
$l = [1, 2, 3, 4, 5]
for $i in $l {
    print $i
}

// While loops are similar:
$i = 0
while $i < 5 {
    print $i
    $i += 1
}

/*
    Finally, functions are somewhat comparable to labels in asm
*/
@info {
    print "Info given here"
    $new_number = $old_number + 10
    return $new_number 
}

$old_number = 5
$number = @info
assert $number, 15

/*
    Multi-line comments are initiated like so
*/

```