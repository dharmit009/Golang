# Modules

Modules are a collection of packages. They are created by having a
go.mod file in the root directory of your project. Module contains the 
information of your project. 

1. All GO projects have a go.mod file.
1. Modules can be managed by GO CLI.
1. Stores information about project which includes:
    - Dependencies.
    - GO version.
    - Package Info.

**Example:**

``` go

module example.com/practice

go 1.17

require(
    github.com/alexflint/go-arg v1.4.2
    github.com/fatih/color v1.13.0
)

```

