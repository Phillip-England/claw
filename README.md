# claw
Bundle up your web components into a single `js` file.

## Installation
```bash
go get github.com/Phillip-England/claw
```

## Usage Rules
Claw expects all the `.js` files in a specific directory to each contain a single class. Also, all web components may only have a two-part name like `NiceButton` or `CoolNav`.

Following these rules enables claw to automatically include all web components on your behalf.

Do not include other Javascript, only web components. Claw is intended only for web components.

# Usage
With the Usage Rule in mind, run:
```go
err := claw.BundleWebComponents("./components", "./static/index.js")
if err != nil {
  fmt.Println(err.Error())
}
```

