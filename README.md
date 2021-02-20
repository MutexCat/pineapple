# pineapple
pineapple programming language demo.
This repository is an implementation of the [pineapple language tutorial](https://github.com/karminski/pineapple)
which is very helpful to Golang learners.
Based on the origin pineapple language,
i also expanded it with math operations
 in a stupid way.

## Usage:
```
go build

./pineapple hello
```

```
SourceCharacter ::=  #x0009 | #x000A | #x000D | [#x0020-#xFFFF]
Name            ::= [_A-Za-z][_0-9A-Za-z]*
StringCharacter ::= SourceCharacter - '"'
String          ::= '"' '"' Ignored | '"' StringCharacter '"' Ignored
VariableStr     ::= "$" Name Ignored
VariableInt     ::= "@" Name Ignored
AssignmentStr   ::= VariableStr Ignored "=" Ignored String Ignored
AssignmentInt   ::= VariableInt Ignored "=" Ignored Int Ignored
Add             ::= "add" "(" Ignored VariableInt Ignored "," Ignored VariableInt Ignored ")" Ignored
Sub             ::= "sub" "(" Ignored VariableInt Ignored "," Ignored VariableInt Ignored ")" Ignored
Mult            ::= "Mult" "(" Ignored VariableInt Ignored "," Ignored VariableInt Ignored ")" Ignored
Div             ::= "Div" "(" Ignored VariableInt Ignored "," Ignored VariableInt Ignored ")" Ignored
Print           ::= "print" "(" Ignored Variable Ignored ")" Ignored
Statement       ::= Print | Assignment | Add | Sub | Mult | Div
SourceCode      ::= Statement+ 
```