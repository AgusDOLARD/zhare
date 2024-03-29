# Zhare

Zhare is a command-line tool designed to simplify file sharing within your local network.

## ✨ Features

- *Simple and Fast:* Quickly share files with others on your local network.
- *QR Code Generation:* Generate a QR code for easy access.
- *No External Dependencies:* Single binary for straightforward use.
- *Cross-Platform:* Works on Windows, macOS, and Linux.

## 🚀 Installation

To install Zhare, use the following command:

```sh
go install github.com/AgusDOLARD/zhare@latest
```

## 📝 Usage

Specify the path to the file you want to share:

```sh
zhare path/to/your/file
```

Or simple run `zhare` to run a server where you can upload files

### 🚩Flags

- _-qr:_ show qr for web page
- _-p:_ server port

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
