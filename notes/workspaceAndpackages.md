# Workspaces:

Programmers typically has one workspace for many different projects.
Golang recommends the following directory hierarchy:

-> 1. src: Contains source code files.
-> 2. pkg: Contains packages (libraries).
-> 3. bin: Contains executables.

This makes extremely easy to share code with other people.

## Packages:

Your code is organized into packages. A package is group of related
source code files. Each package can be imported by other packages which
is very useful for software reuse.

1. There should be always one package named `main`.
1. Building the main package generates an executable program.
