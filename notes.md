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
- on parseProgram (p36)
    - inits an AST node, starts moving through tokens
    - starts an if with null, if we're looking at a let, return, if then we call parsing functions for those
        - the parse functions check the identifier after the let statement, check for an equals (if not return null with errormsg), check the net token, then look at the expression
        - we create a variable statement node and then add on the iden and value, then return the variable statement
    - we then push the net statement onto the AST and advance tokens, then return program
- so we're basically encoding the semantics of our language here, such as erroring out unless we're following rules like 'let <ident> = expression', building this part of the tree, then moving deeper
    - for expression, we're doing the same

### 2.7 - how Pratt Parsing works

