# Pastiche

A clipboard code formatting daemon with a lovely extendable `Formatter` interface. Copy source code from a PDF file or agentic coding tool, hit your hotkey (or just copy and wait), and `pastiche` formats it in place — gofmt, black, rustfmt, or anything you plug in.

## Motivation

We as programmers have to constantly improve our skills, and in my case that means reading a whole lotta books. Specifically reading books in PDF(yes, I am a rare connaisseur of eBooks) and then saving important information in the Obsidian vault that I maintain. And copying code from sources that do not preserve formatting has always been a hassle, though with code editors and IDEs nowadays providing automatic formatting and linting on save or hotkey press life became much easier. But the problem still exists, when there is no formatter available at hand's reach -- pressing tabs and spaces inside backtikcs gets painful really fast. This is why `pastiche` exists, to make sure that the code is formatted before it even hits the document. 

## Quick Start

Before running the daemon, install it. Make sure that your have a Go toolchain available on your system, otherwise please see [this page](https://golang.org/doc/install) first.

```sh
go get github.com/inflame-ue/pastiche@latest
```

You will then be able to configure the application and either run it in a blocking termainl window or install it as a user-level daemon:

```sh
# Configure the trigger, hotkey, etc
pastiche configure

# Run in a terminal window
pastiche

# Install as a user-level daemon
pastiche install
```

**Note:** If no config file exists, `pastiche daemon` and `pastiche status` will fail with an error. Therefore, make sure to always run `pastiche configure`, which will create a `pastiche.toml` configuration file in your `.config` directory.

## Usage

### Trigger modes

| Mode | Behavior |
|---|---|
| `autowatch` | Detects code in clipboard changes and formats automatically |
| `hotkey`   | Waits for a keypress (e.g. Ctrl+I) to read clipboard and format |
| `both`     | Autowatch in background + hotkey for man
ual formatting |

### Built-In Formatters

| Language | Tool | Status |
|---|---|---|
| Go | `go/format`  | Built-in, no external dep |
| Python | `black` | Requires `pip install black` |
| Rust | `rustfmt` | Requires `rustup component add rustfmt` |


### Commands

```
pastiche            Run the daemon (default)
pastiche configure  Open the TUI config editor
pastiche install    Install as a systemd user service
pastiche status     Show config and formatter availability
```

### Config

Written to `~/.config/pastiche/pastiche.toml`. The `configure` command edits it interactively.

```toml
[trigger]
mode = "autowatch"

[hotkey]
key = 9

[heuristic]
value = 3

[formatters]
order = ["go", "python", "rust"]
```


## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for development setup instructions.



