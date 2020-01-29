Testing
==

Each test has a name, which must be listed in the file `tests.list`
and may not contain funny characters.

Running the tests is done with `make test` or directly with:

```
$ ./run-tests
```

If any changes are reported, you should either fix the program so as
to fix the output, or accept the output.

Accepting the output is done by replacing the `*.expected` files. It
can be done after running the tests in one fell swoop by running:

```
$ ./accept-changes
```
