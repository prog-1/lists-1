# Singly Linked List

## Useful Links

1. https://goplay.tools/snippet/_ByxpX_E8Ll
2. https://en.wikipedia.org/wiki/Linked_list

## Modules & Workspaces

This section explains some Go-specific management features used in this repo.

Package `github.com/prog-1/list` is implemented as a separate subpackage. Go
provides an easy way to work with multiple packages. More information available
at https://go.dev/doc/tutorial/workspaces.

In addition to that, we replace `github.com/prog-1/list` module paths to ensure
this repo is self-contained and no external repos or packages are required. More
information available at at https://go.dev/ref/mod#go-mod-file-replace.

## Home Exercises

In the exercises below you will have to update only those files, that already exist
in this repo. You must not add new files, methods or functions unless you need them
for testing or implementation details.

You have to add tests only for those functions that you have to implement or modify
in the exercises below.

1. Implement missing `list.List` methods and tests.
2. Implement `InsertionSort` and `SortedInsert` functions and tests.

You can earn additional points (up to 50%) for `InsertionSort` implementation which
swaps the element pointers without destroying and creating elements. Please note that
this may require additional `List` methods, since you cannot modify `Element.next` field
outside of `github.com/prog-1/list` package (and making the field public is not desired).
