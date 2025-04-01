# gon

> Scaffold models, usecases, and handlers for Go apps with an opinionated directory layout — just like Rails, but for Go.

---

## ✨ Rails-like Developer Experience

Inspired by the Rails philosophy of *convention over configuration*, `gon` lets you scaffold everything you need for a domain entity — model, repository, usecase, and handler — with a single command.

```bash
gon g scaffold User name:string email:string
```

This generates fully structured code under `internal/domain/user/`, just like `rails g scaffold` — but in idiomatic Go.

---

## ✨ Features

- Rails-style generators for Go projects
- Generate boilerplate code for Clean Architecture
- Support for models, repositories, usecases, and handlers
- Enforces consistent directory structure
- Lightweight and fast CLI

---

## 🚀 Installation

### Go 1.17 or later

```bash
go install github.com/mickamy/gon@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Go 1.24 or later (with `go get -tool`)

```bash
go get -tool github.com/mickamy/gon
```

This installs `gon` to `$GOTOOLDIR/bin` (usually `$HOME/go/bin`).

---

## 📦 Initial Setup

### Step 1: Create configuration file

```bash
gon init
```

This creates a `gon.yaml` file with default settings. You can tweak output paths, package names, and template locations.

### Step 2: Install templates and dependencies

```bash
gon install
```

This command generates the database file and prepares templates required for scaffolding, including:

- Embedded templates for model, repository, usecase, handler
- Includes lightweight test helpers like `httptestutil/request.go`, generated during install for better testing experience.

> 💡 Make sure to run this before using `gon g` or `gon d`.

---

## 🧪 Usage

### Generate a domain model

```bash
gon generate model User name:string email:string
# or simply
gon g model User name:string email:string
```

### Generate a usecase

```bash
gon g usecase CreateUser
```

### Generate a handler

```bash
gon g handler User list create
```

### Scaffold everything

```bash
gon g scaffold User name:string email:string
```

This generates model, repository, usecase, handler, and fixture in one shot.

### Destroy everything

```bash
gon d scaffold User
```

This deletes generated files for the given domain entity.

---

## 📁 Output Structure

```text
internal/
└── domain/
    └── user/
        └── fixture/
        │   └── user.go
        ├── model/
        │   └── user_model.go
        ├── usecase/
        │   ├── create_user_use_case.go
        │   ├── get_user_use_case.go
        │   ├── list_user_use_case.go
        │   ├── update_user_use_case.go
        │   └── delete_user_use_case.go
        ├── repository/
        │   └── user_repository.go
        └── handler/
            └── user_handler.go
test/
└── httptestutil/
    └── request.go
```

> Each subdirectory under `domain/{name}` is a separate package.

---

## 🧪 Testing & Fixtures

When generating code, `gon` also prepares test helpers to improve DX:

- `httptestutil.RequestBuilder` to build and test Echo requests
- Zero-value based fixtures with `TODO` comments for easy customization

```go
package fixture

func User(setter func(*model.User)) model.User {
  m := model.User{
    // TODO: fill in default values
  }
  if setter != nil {
    setter(&m)
  }
  return m
}
```

---

## 🛠 Template Driven

Templates are embedded using Go 1.16+ `embed` package. You can customize them by copying from the embedded defaults
during `gon install`.

---

## 📚 Example Project

You can find a working example project using `gon` under the [`example/`](./example) directory.

This example demonstrates how the generated code looks and how to structure your application using `gon`'s opinionated layout. It’s a great starting point to explore and adapt into your own Go project.

---

## 📄 License

[MIT](./LICENSE)
