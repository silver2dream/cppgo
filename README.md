# cppgo

This library allows methods on C++ objects to be called directly from the
Go runtime without requiring cgo compilation.

To set up a Go object that proxies method calls to a C++ object:

1. Define a Go struct type with function pointer field declarations that match
   the C++ class declaration,
2. Get the address of the C++ object in memory as a `uintptr` type,
3. Call `cpp.ConvertRef(addr, &o)` to point this proxy struct object (`o`) at
   the C++ object by its `addr`.
4. After this initial setup, the function pointers on the struct object will be
   ready to use like any Go method.

## Example Usage

The following example will call into a C++ class `Library` with a
`GetString(char *name)` object method prototype. See the "STEP X" comments
for usage guides:

```go
package main

var (
  dll = syscall.MustLoadLibrary("mylib.dll")
  create = dll.MustFindProc("new_object@0") // @0 due to stdcall mangling
)

// STEP 1. define our C++ proxy struct type with function pointers.
// The functions in this struct will be filled in by the `cpp.ConvertRef()`
// function, at which point this struct will proxy all method calls to the
// C++ object.
type Library struct {
  GetString(name string) string
}

func main() {
  // STEP 2. get an address for the C++ object
  // NOTE: you may need to free this later depending on call semantics.
  o, _, _ := create.Call() // o is a uintptr

  // STEP 3. point our proxy structure at the functions located in the object
  // that we got from step 2.
  if err := cpp.ConvertRef(o, &l); err != nil {
    panic(err)
  }

  // STEP 4. call the function with arguments
  fmt.Println(l.GetString("Loren"))
}

// Prints:
// Hello, Loren!
```

The C++ class for the above program could look something like:

```cpp
#include <stdio.h>

#ifndef WIN32
#  define __stdcall
#  define __cdecl
#  define __declspec(x)
#endif

class Library {
public:
  virtual char __cdecl *GetString(char *name) {
    return sprintf("Hello, %s!", name);
  }
}

extern "C" __declspec(dllexport) Library* __stdcall new_object() {
  return new Library();
}
```

## Using `__stdcall` Calling Convention on Windows

By default, this library expects that methods will use the `__cdecl` calling
convention, which is the default for C/C++ POSIX compilers (Windows uses
`__thiscall` which is not supported, see below for caveats). If you want to
opt into the `stdcall` calling convention, you can do so with a `call:"std"`
tag on the field declaration:

```go
type Library struct {
  GetID() int `call:"std"` // "stdcall" is also valid here
}
```

This will ensure that the function pointer is compatible with the library's
calling convention.

**Note**: The `__thiscall` calling convention (default convention for C++
method functions) is unsupported in this library. See the caveats section
for more information.

## Caveats & Gotchas

### Limited Type Conversion

You can define arbitrary function arguments and return types, but the internal
translation does not support a full range of types. Currently, only the
following types are supported:

* `uint*`, `int*`, `string`, `uintptr`

Any other type of pointer will be passed as a pointer directly, which may be
what you want, but may not be. For any type that isn't well converted by
the library, use `uintptr` to send its address.

Note that slices are not well supported due to the extra information
encoded in a Go slice.

Note also that `string` converts only to and from the `char*` C type, in other
words, C strings. The `std::cstring` or `wchar_t` types are not yet supported.

### Passing Objects as Arguments

When passing C++ objects to methods, you will want to use the `cpp.Ptr` or
`uintptr` value representing the address of the object. You cannot use the
pointer to the proxy struct, since this does not actually point to the
object, and cppgo does not know how to translate between the two.

For example, to send objects to methods, define a function and call it using
the bare address pointers:

```go
// C++ class defined as:
//
// class MyClass {
//   virtual void DoWork(MyClass *other);
// }
type MyClass struct {
  DoWork(obj uintptr)
}

func main() {
  var m1 MyClass
  var m2 MyClass
  o1 := get_object_address() // uintptr
  o2 := get_object_address() // uintptr

  cpp.ConvertRef(o, &m1)
  cpp.ConvertRef(o, &m2)

  // we may have m2 here, but we call DoWork() with the uintptr address.
  m1.DoWork(o2)
}
```

### Non-Virtual Functions Not Supported

This library does not yet support non-virtual functions. Only functions
defined with the `virtual` keyword are callable.

### No `__thiscall` Support on Windows

By default, Windows compilers use `__thiscall` calling conventions for
C++ methods. This calling convention is, unfortunately, incompatible with
Go's syscall support, and is not supported. Note that `__cdecl` (the default
C/C++ calling convention) _is_ supported on Windows, and you can happily
annotate your functions with that call convention.

## Author & License

Written by Loren Segal in 2017, licensed under MIT License (see LICENSE).
