Python is an object-oriented programming language, and in Python everything is an object.

Let's flesh-out what this means. Earlier we saw that variables are simply
pointers, and the variable names themselves have no attached type information.
This leads some to claim erroneously that Python is a type-free language. But
this is not the case! 

Consider the following code on the right.

Python has types; however, the types are linked not to the variable names but
to the objects themselves.

Lists are ordered collections of items. They can contain any type of data and are created using square brackets `[]`.

Common list operations:

- Create a list with `[]`
- Add items with `append()`
- Access items with index `[0]`
- Get length with `len()`
- Slice lists with `[start:end]`

Try this example:
