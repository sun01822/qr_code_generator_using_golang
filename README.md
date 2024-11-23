
# QR Code Generator in Go

This is a simple Go application that generates a QR code for a specified URL and saves it as a PNG file. The project uses the [`go-qrcode`](https://github.com/skip2/go-qrcode) library.

---

## Features

- Generates a QR code for any URL.
- Saves the QR code as a PNG image.
- Easy to customize and extend.

---

## Requirements

- Go installed on your machine. Download it from [golang.org](https://golang.org/dl/).
- The `go-qrcode` library (automatically installed via the `go get` command).

---

## Installation

1. Clone this repository or create a new directory for the project:

   ```bash
   git clone https://github.com/sun01822/qr_code_generator_using_golang.git
   cd qr_code_generator_go
   ```

2. Install the required library:

   ```bash
   go get -u github.com/skip2/go-qrcode
   ```

3. Build and run the application:

   ```bash
   go run main.go
   ```

---

## Usage

1. Open the `main.go` file.
2. Replace the URL `"http://github.com/sun01822"` with the URL for which you want to generate a QR code.
3. Run the application using:

   ```bash
   go run main.go
   ```

4. The QR code will be saved as `qrcode.png` in the same directory as the application.

---

## Example

By default, this application generates a QR code for the URL `http://github.com/sun01822` and saves it as `qrcode.png`.

Generated QR code file:  
![Example QR Code](./qrcode.png)

---

## Customization

### Change URL

Replace the URL string in the `main.go` file with your desired URL:

```go
err := qrcode.WriteFile("http://example.com", qrcode.High, 256, "qrcode.png")
```

### Adjust Image Size

Modify the size parameter (e.g., `256`) to change the dimensions of the QR code:

```go
err := qrcode.WriteFile("http://example.com", qrcode.High, 512, "qrcode.png")
```

---

## Libraries Used

- [`skip2/go-qrcode`](https://github.com/skip2/go-qrcode): A Go library for generating QR codes.

---

## License

This project is open-source and available under the MIT License.

---

## Author

Developed by [MD. Shariar Hossain Sun](https://github.com/sun01822).  
Feel free to reach out for contributions or issues!
