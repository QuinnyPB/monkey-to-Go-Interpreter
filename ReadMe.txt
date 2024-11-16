What can this interpreter do?:
handles C-like syntax,
variable binding,
integers and bools,
arithmetic expressions,
prefix and infix operators,
built-in functions,
first-class and higher-order functions,
closures,
string, array and hash data structures,
processing macros and their expansion,

Major components of the interpreter:
A Lexer
A Parser
An Abstract Syntax Tree (AST)
An Internal Object System
An Evaluator

Implements Pratt Parsing (Top-Down Operator Precedence) for the AST. This handles the logic for giving precedence to equations and expressions. E.g. (4 + 1) * 3 - 2 -> (((4 + 1) * 3) - 2)

Is a "tree-walking interpreter", which takes the AST built by the parser and interprets it on the go, without pre-processing or compilation.

Challenges for Myself:
Allow the lexer to fully support Unicode, such as emojis and other characters.
Allow lexer to handle floats, hex notation, octal notation