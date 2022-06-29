# Init-less Cobra

An example of a cobra cli without using init and encapsulating command creation.

## Goals

* No init functions
* Minimal `main()` function
* Minimal behavior in commands
  * Get needed vars and call to internal package to perform work
  * Do **not** include business logic
* Dependency injection for internal worker functions:
  * Only require what's actually needed
  * Limits function's requirements and makes testing easier

## Project Structure

```text
.
├── cmd
│  ├── create.go
│  └── root.go
├── go.mod
├── go.sum
├── internal
│  └── accounts
│     ├── client.go
│     └── users
│        └── users.go
└── main.go
```

## Usage

```text
╭─16:23 gemini ~/code/mxygem/cobrainitless ‹main*›
╰─> go build -o ./bin/accounter
╭─16:23 gemini ~/code/mxygem/cobrainitless ‹main*›
╰─> ./bin/accounter
a tool to manage various aspects and behaviors relating to accounts

Usage:
  accounter [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      create a new user
  help        Help about any command

Flags:
  -d, --debug   enable debugging
  -h, --help    help for accounter

Use "accounter [command] --help" for more information about a command.
╭─16:23 gemini ~/code/mxygem/cobrainitless ‹main*›
╰─> ./bin/accounter create
2022/06/29 16:23:23 create command sending email to users.Create: ""
Error: cannot create user, no email provided
Usage:
  accounter create [flags]

Flags:
  -e, --email string   specify an email to use
  -h, --help           help for create
  -t, --type string    the resource type to create. options: [user] (default "user")

Global Flags:
  -d, --debug   enable debugging

╭─16:23 gemini ~/code/mxygem/cobrainitless ‹main*›
╰─> ./bin/accounter create -t organization
Error: unknown resource type of "organization" received
Usage:
...

╭─16:23 gemini ~/code/mxygem/cobrainitless ‹main*›
╰─> ./bin/accounter create -t user
2022/06/29 16:23:29 create command sending email to users.Create: ""
Error: cannot create user, no email provided
Usage:
...

╭─16:23 gemini ~/code/mxygem/cobrainitless ‹main*›
╰─> ./bin/accounter create -t user -e foo@bar.com
2022/06/29 16:23:32 create command sending email to users.Create: "foo@bar.com"
2022/06/29 16:23:32 creating user with email: "foo@bar.com"
```
