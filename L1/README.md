# Go Basics Notes

## 1. Print to Console

```go
fmt.Println("Hi Mohammed")
```

* Function from the `fmt` package
* Used to print text to the console

---

## 2. Entry Point (main function)

```go
package main

func main() {
    // code here
}
```

* The program must use `package main`
* It must include `func main()`
* This is the entry point of any Go program

---

## 3. Case Sensitivity

* Go is case-sensitive
* `main` is not the same as `Main`

---

## 4. Run Code

```bash
go run main.go
```

What happens:

1. Compile
2. Execute

---

## 5. Build Code

```bash
go build
```

* Compiles code into an executable file
* Faster execution afterward because it is already compiled

---

## 6. How Go Executes Code

* Go is not interpreted
* Code is compiled into **machine code**
* The CPU executes machine code directly

---

## 7. Output Files by OS

* Windows → `.exe`
* Linux / macOS → binary file (no extension)

---

## 8. Example

```go
package main

import "fmt"

func main() {
	fmt.Println("Hi Mohammed")
}
```

Run:

```bash
go run main.go
```

---

## Notes

* Execution starts only from `main()`
* You can call other functions from `main()`
* Go produces a static binary (no runtime like Node.js or Python)
