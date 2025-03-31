# gon

> Scaffold models, usecases, and handlers for Go apps with an opinionated directory layout.
>

---

## ✨ Features

- Generate boilerplate code for Clean Architecture
- Support for models, repositories, usecases, and handlers
- Enforces consistent directory structure
- Lightweight and fast CLI

---

## 🚀 Installation

```
go install github.com/mickamy/gon@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

---

## 🧪 Usage

### Generate a domain model

```
gon generate model User name:string email:string
```

### Generate a usecase

```
gon generate usecase CreateUser
```

### Generate a handler

```
gon generate handler UserHandler --with-usecase
```

---

## 📁 Output Structure

```
internal/
├── domain/
│   ├── model/
│   │   └── user.go
│   ├── usecase/
│   │   └── create_user.go
│   ├── repository/
│   │   └── user_repository.go
│   └── handler/
│       └── user_handler.go
```

---

## 🛠 Template Driven

Templates are embedded using Go 1.16+ `embed` package. You can customize them later.

---

## 📄 License

[MIT](./LICENSE)
