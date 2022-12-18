<p align="center"><img alt="pix-utils" src="./assets/logo-pix.png" width="128px" /></p>

# <p align="center">Go-Pix-Utils<p>

<p align="center">
  <a
    href="https://github.com/thalesog/go-pix-utils/blob/master/LICENSE"
    target="blank"
  >
    <img
      src="https://img.shields.io/github/license/thalesog/go-pix-utils"
      alt="License"
    />
  </a>
  <a href="https://github.com/thalesog/go-pix-utils/stargazers" target="blank">
    <img
      src="https://img.shields.io/github/stars/thalesog/go-pix-utils"
      alt="Stars"
    />
  </a>
<a href="https://pkg.go.dev/github.com/thalesog/go-pix-utils/pixUtils" target="blank">
    <img
      src="https://pkg.go.dev/badge/github.com/thalesog/go-pix-utils"
      alt="GoDoc"
    />
  </a>
<a href="https://goreportcard.com/report/github.com/thalesog/go-pix-utils" target="blank">
    <img
      src="https://goreportcard.com/badge/github.com/thalesog/go-pix-utils?force=true"
      alt="Go Report Card"
    />
  </a>
</p>

> ### ⚠️ This package is under development and is not ready for production

> Pix-Utils is a set of tools to parse, generate and validate payments of Brazil Instant Payment System (Pix), making fast and simple to handle charges and proccess then in your project. Originally developed using [TypeScript](https://github.com/thalesog/pix-utils), this is the Go version of the library.

# 🚀 Features

- [x] Parse Static Pix EMV
- [x] Parse Dynamic Pix EMV
- [x] Validate CRC16
- [x] Generate Static Pix EMV
- [x] Generate Dynamic Pix EMV
- [ ] Generate QRCode Image from EMV or Pix

This library also normalize the input data to the Pix standard, so you don't need to worry about it.
It follows the [Pix Specification](https://www.bcb.gov.br/content/estabilidadefinanceira/pix/Regulamento_Pix/II_ManualdePadroesparaIniciacaodoPix.pdf), that mainly contains the following guidelines about EMV data structure:

- Merchant Name: 25 characters
- Merchant City: 15 characters
- Pix Key/Payload URL: 77 characters

# 📦 Installation

### Install the package in your project

```sh
go get github.com/thalesog/go-pix-utils
```

### Create Static Pix

```go
package main

import (
  "fmt"
  "github.com/thalesog/go-pix-utils/pixUtils"
)

func main() {
  pix := pixUtils.CreateStaticPix(pixUtils.CreateStaticPixParams{
    MerchantName:      "Thales Ogliari",
    MerchantCity:      "São Miguel do Oeste",
    PixKey:            "thalesog@me.com",
    TransactionAmount: 10.00,
    AditionalData:     "Pedido 123",
  })

  fmt.Printf("Pix Type: %s \n", pix.Type)
  fmt.Printf("EMV Code: %s \n", pix.Raw)
  fmt.Printf("Pix Elements: %v \n", pix.Elements)
}
```

### Create Dynamic Pix

```go
package main

import (
  "fmt"
  "github.com/thalesog/go-pix-utils/pixUtils"
)

func main() {
  pix := pixUtils.CreateDynamicPix(pixUtils.CreateDynamicPixParams{
    MerchantName: "Thales Ogliari",
    MerchantCity: "São Miguel do Oeste",
    Url:          "https://pix.thalesogliari.com.br",
  })

  fmt.Printf("Pix Type: %s \n", pix.Type)
  fmt.Printf("EMV Code: %s \n", pix.Raw)
  fmt.Printf("Pix Elements: %v \n", pix.Elements)
}
```

# 🍰 Contributing

Please contribute using [GitHub Flow](https://guides.github.com/introduction/flow). Create a branch, add commits, and [open a pull request](https://github.com/thalesog/go-pix-utils/compare).

# 📝 License

This project is under [MIT](https://github.com/thalesog/go-pix-utils/blob/master/LICENSE) license.

#

<p align="center">
 Developed with 💚 by <a href="https://github.com/thalesog">@thalesog</a> 🇧🇷
</p>
