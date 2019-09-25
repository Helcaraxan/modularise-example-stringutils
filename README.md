# Modularise - Demo project

This repository contains a piece of a demo setup for the _modularise_ tool. The proposal and
technical design document for this tool can be found
[here](https://docs.google.com/document/d/1g7wIlxn9JBJkc-2lHAb2MfCYLltWZKyNlIO6SBW1x-8/edit#).

## Project layout

The project implements a placeholder command-line implementation. The command-line is capable of
printing random numbers and strings based on being passed either "number" or "string as argument.
An evolution of the tool also supports passing a second length argument in the case of a first
"string" parameter.

The project's implementation is spread across four separate packages:

1. The `main` package that implements the CLI.
2. The `numberutils` package that exposes the functionality of printing a random number.
3. The `stringutils` package that exposes the functionality of printing a random string.
4. The `random` package that exposes the generation of random numbers and strings.

For the demonstration's purposes all packages should be considered as implementing related but
standalone APIs. The overall project's end-product is the CLI tool itself but downstream users
may also depend on the underlying implementation APIs to satisfy advanced use-cases.

## Repository layout

- The "core" repository can be found at [Helcaraxan/modularise-example-core]. It contains four
  separate branches:
  1. A `master` branch that contains an initial version of the CLI tool.
  2. A `master-extension` branch that contains the CLI tool with the length argument added to
    string printing.
  3. A `modularised` branch that contains a "modularised" version of `master`.
  4. A `modularised-extension` branch that contains a "modularised" version of `master-extension`.
- The "numberutils" split can be found at [Helcaraxan/modularise-example-numberutils]. It contains
  the extracted source code of the `numberutils` project of the "core" repository. To satisfy its
  dependencies it also contains a copy of the `random` package in its `internal/` subtree. It
  contains one branch:
  1. A `master` branch that contains the extracted code of the `numberutils` package from the "core"
    repository's `modularised` branch.
- The "stringutils" split can be found at [Helcaraxan/modularise-example-stringutils]. It contains
  the extracted source code of the `stringutils` project of the "core" repository. To satisfy its
  dependencies it also contains a copy of the `random` package in its `internal/` subtree. It
  contains two branches:
  1. A `master` branch that contains the extracted code of the `stringutils` package from the "core"
    repository's `modularised` branch.
  2. A `master-v2` branch that contains the extracted code of the `stringutils` package from the
    "core" repository's `modularised-extension` branch.

## Walk-through of the modularise tool usage

The `master` and `master-extension` branches in the "core" repository represent the evolution of
the project without the use of the _modularise_ tool. Adding the ability to specify a string length
is a non-breaking change from the CLI perspective and, if a new release were to be tagged, it would
be considered a new minor version of the project's system-level API. However from a Go module
perspective it should be a major release given that there has been a breaking change in the
`stringutils` API.

In the case of the "modularised" project, the `numberutils` and `stringutils`, as well as the
`random` implementation-level APIs have been moved to an `internal/` subtree. The `numberutils` and
`stringutils` are both mirrored in the split repositories.

As a result:

- The `modularised-extension` branch could be released as a "proper" new module minor version of the
  project, contrary to the `master-extension` branch, because there is no breaking change in the
  core repository's public API.
- The `modularised-extension` branch corresponds, like the `modularised` branch, to the HEAD commit
  on the `master` branch in the numberutils repository because there are no changes to the
  `numberutils` package in the core repository.
- The `modularised-extension` branch corresponds to the HEAD commit of the `master-v2` branch in
  the stringutils repository, unlike the `modularised` branch which maps to the HEAD commit of the
  repository's `master` branch. The `master-v2` branch will be released as a new major version of
  the split because of the breaking change to its public API.

Summary of the branch mappings between the core and split repositories:

| core                    | numberutils | stringutils |
| ----------------------- | ----------- | ----------- |
| `master`                | _N.A_       | _N.A_       |
| `master-extension`      | _N.A_       | _N.A_       |
| `modularised`           | `master`    | `master`    |
| `modularised-extension` | `master`    | `master-v2` |

[Helcaraxan/modularise-example-core]: https://github.com/Helcaraxan/modularise-example-core
[Helcaraxan/modularise-example-numberutils]: https://github.com/Helcaraxan/modularise-example-numberutils
[Helcaraxan/modularise-example-stringutils]: https://github.com/Helcaraxan/modularise-example-stringutils
