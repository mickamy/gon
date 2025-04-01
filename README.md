# gon

> Scaffold models, usecases, and handlers for Go apps with an opinionated directory layout.

---

## ✨ Features

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
gon g handler User
```

---

## 📁 Output Structure

```text
internal/
└── domain/
    └── user/
        ├── model/
        │   └── user_model.go
        ├── usecase/
        │   └── create_user_use_case.go
        ├── repository/
        │   └── user_repository.go
        └── handler/
            └── user_handler.go
```

> Each subdirectory under `domain/{name}` is a separate package.

---

## 🛠 Template Driven

Templates are embedded using Go 1.16+ `embed` package. You can customize them by copying from the embedded defaults
during `gon install`.

---

## 📄 License

[MIT](./LICENSE)
