# All Training Articles:
1. [Git Config](https://stackoverflow.com/questions/4220416/can-i-specify-multiple-users-for-myself-in-gitconfig)
  * [Git Branching Arch](https://nvie.com/posts/a-successful-git-branching-model/)
  * git log --pretty --oneline --graph
  * [Try Git Visually Tutorial](https://try.github.io/)
  * git config --global core.editor "code --wait"/'vi'
  * [My Git Notes](https://docs.google.com/document/d/1HESVXebz3x3i1RD7Tyh8k1K82dNKeaWNs4tPQ0oZD-w/edit)

2. MYSQL Links:
  * [MySQL Privilages](https://linuxize.com/post/how-to-create-mysql-user-accounts-and-grant-privileges/)
  * [PackageSQL](https://golang.org/pkg/database/sql/)
  * [QueryRowContext](https://golang.org/pkg/database/sql/#DB.QueryRowContext)
  * [Query](https://golang.org/pkg/database/sql/#DB.Query)
  * [SQL Mock](https://github.com/DATA-DOG/go-sqlmock)
  * [Example Orders for Mock](https://github.com/DATA-DOG/go-sqlmock/blob/master/examples/orders/orders_test.go)
  * [Truncate a table referenced by a foreign key](https://stackoverflow.com/questions/5452760/how-to-truncate-a-foreign-key-constrained-table)
  
3. Gorilla Mux:
  * [Mux Doc](https://pkg.go.dev/github.com/gorilla/mux)
  * Example: r.HandleFunc("/products", ProductsHandler).
    Host("www.example.com").
    Methods("GET").
    Schemes("http")
  * [Mux Examples](https://github.com/gorilla/mux#examples)
  * [Mux SetURLVars](https://pkg.go.dev/github.com/gorilla/mux#SetURLVars)

4. golangci-lint:
  * [Install](https://golangci-lint.run/usage/install/#linux-and-windows)
  * Set up your own linters: Sample golangci-lint file
  ```
linters-settings:
  govet:
    check-shadowing: true
  golint:
    min-confidence: 0
  gocyclo:
    min-complexity: 10
  gocognit:
    min-complexity: 10
  maligned:
    suggest-new: true
  dupl:
    threshold: 100
  goconst:
    min-len: 2
    min-occurrences: 3
  misspell:
    locale: US
  lll:
    line-length: 140
  goimports:
    local-prefixes: gitlab.kroger.com/krogo
  gocritic:
    enabled-tags:
      - diagnostic
      - experimental
      - opinionated
      - performance
      - style
    disabled-checks:
      - wrapperFunc
      - dupImport # https://github.com/go-critic/go-critic/issues/845
      - ifElseChain
      - octalLiteral
  funlen:
    lines: 100
    statements: 50

linters:
  # please, do not use `enable-all`: it's deprecated and will be removed soon.
  # inverted configuration with `enable-all` and `disable` is not scalable during updates of golangci-lint
  disable-all: true
  enable:
    - asciicheck
    - bodyclose
    - deadcode
    - dogsled
    - dupl
    - errcheck
    - exportloopref
    - exhaustive
    - funlen
    - gochecknoglobals
    - gochecknoinits
    - gocognit
    - goconst
    - gocritic
    - gocyclo
    - godox
    - gofmt
    - goimports
    - golint
    - gomnd
    - gosec
    - gosimple
    - govet
    - ineffassign
    - lll
    - maligned
    - misspell
    - nakedret
    - nolintlint
    - prealloc
    - rowserrcheck
    - scopelint
    - sqlclosecheck
    - staticcheck
    - structcheck
    - stylecheck
    - typecheck
    - unconvert
    - unparam
    - unused
    - varcheck
    - whitespace
    - wsl

# golangci.com configuration
# https://github.com/golangci/golangci/wiki/Configuration
service:
  golangci-lint-version: 1.28.x # use the fixed version to not introduce new linters unexpectedly
  prepare:
    - echo "here I can run custom commands, but no preparation needed for this repo"
  ```

5. PostMan Tests:
* [Test-Scripts](https://learning.postman.com/docs/writing-scripts/test-scripts/)
* [Testing Status Codes](https://learning.postman.com/docs/writing-scripts/script-references/test-examples/#testing-status-codes)
* [Testing Response Body](https://learning.postman.com/docs/writing-scripts/script-references/test-examples/#testing-response-body)

6. GoLang Technology Essentials ZopSmart:
* [Read it](https://docs.zopsmart.com/doc/technology-essentials-9RZ5dcHfel)
* [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments#variable-names)
* [Golang Design Patterns](https://golangbyexample.com/golang-comprehensive-tutorial/)

7. Handling Errors:
* [Writing constant errors](https://medium.com/@smyrman/writing-constant-errors-with-go-1-13-10c4191617)

8. Coverage Reports:
* [All commands](https://blog.seriesci.com/how-to-measure-code-coverage-in-go/)
