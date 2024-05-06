# notes on the book

- building a tree-walking interpreter
    - parses source code, builds an AST, then evaluates it
- unlike non-evaluating 'bytecode compilers'
- unlike JIT -> native machine code interpreters

---

## extensions

- extend isLetter to expand allowed identifiers
- extend readNumber to support for float & hex

---

## Parsing

- Vaughan Pratt - recursive descent parsing strategy