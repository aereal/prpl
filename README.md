[![status][ci-status-badge]][ci-status]
[![PkgGoDev][pkg-go-dev-badge]][pkg-go-dev]

# prpl

prpl = parameters pull tool

prpl is a tool running command with parameters that stored in [AWS SSM Parameter Store][aws-ssm-param-store].

The parameters are exported as environment variables.

## Synopsis

If you have parameters such as:

- `/my-app/staging/creds/id`
- `/my-app/staging/creds/password`

then run command below and get a result:

```sh
prpl -path /my-app/staging env
# CREDS_ID=<ID>
# CREDS_PASSWORD=<PASSWORD>
```

Environment variable named in below rules:

- Remove `-path` value from full parameter path
  - prpl considers `-path` as a prefix and parameters can be unique without common prefix
  - environment variables names should not have environment name (such as `staging`) for convinience
    - parameters typically have environment in prefix
    - the app may refers environment variables such as `CREDS_ID` not `MY_APP_STAGING_CREDS_ID`
- Replace all characters except for alphabets or numbers with underscore (`_`)
- Convert characters to upper cases

## Installation

```sh
go install github.com/aereal/prpl/cmd/prpl
```

## Motivation

prpl is largely inspired by [ssmwrap][ssmwrap].

prpl have less options to take ease of use.

## License

See LICENSE file.

[pkg-go-dev]: https://pkg.go.dev/github.com/aereal/prpl
[pkg-go-dev-badge]: https://pkg.go.dev/badge/aereal/prpl
[ci-status-badge]: https://github.com/aereal/prpl/workflows/CI/badge.svg?branch=main
[ci-status]: https://github.com/aereal/prpl/actions/workflows/CI
[aws-ssm-param-store]: https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html
[ssmwrap]: https://github.com/handlename/ssmwrap
