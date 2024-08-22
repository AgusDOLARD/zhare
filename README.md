# Zhare

Zhare is a command-line tool designed to simplify file sharing within your local network.

## ✨ Features

- _Simple and Fast:_ Quickly share files with others on your local network.
- _No External Dependencies:_ Single binary for straightforward use.
- _Cross-Platform:_ Works on Windows, macOS, and Linux.

## 🚀 Installation

To install Zhare, use the following command:

```sh
go install github.com/AgusDOLARD/zhare@latest
```

## 📝 Usage

Specify the path to the file you want to share:

```sh
zhare path/to/your/file other/file ...
```

Or, if you want to share a directory, use the `--dir` flag:

```sh
zhare --dir path/to/your/directory
```

### 🚩Flags

- _-p, --port:_ server port
- _-d, --dir:_ directory to serve
- _-v, --version:_ print version
- _--[no]-log:_ enable/disable logging

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
