/*
Package syntax parses regular expressions into parse trees and compiles
parse trees into programs.

Syntax

This package supports following syntax in addition to the golang built-in regexp.


Grouping:
  (?>re)         atomic group; non-capturing
  (?=re)         lookahead; non-capturing
  (?!re)         negative lookahead; non-capturing
  (?<=re)        lookbehind; non-capturing
  (?<!re)        negative lookbehind; non-capturing
  (?{func})      function call; non-capturing
  (?#comment)    comment

Repetitions:
  x*+            zero or more x, possessive
  x++            one or more x, possessive
  x?+            zero or one x, possessive
  x{n,m}+        n or n+1 or ... or m x, possessive
  x{n,}+         n or more x, possessive
  x{n}+          exactly n x, possessive

Back reference:
  \kN            refer to numbered capturing
  \kName         refer to named capturing
  \k{N}          refer to numbered capturing
  \k{Name}       refer to named capturing


Lookbehind limitations

Lookbehind and negative lookbehind only support expressions
that have deterministic matching length.


  (?<=abc)        // OK
  (?<=.{2,5})     // OK
  (?<=foo|barbaz) // OK
  (?<=x+)         // NG

*/
package syntax
