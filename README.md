# PicoNote

[![build](https://github.com/lxsavage/pico-note/actions/workflows/build.yml/badge.svg)](https://github.com/lxsavage/pico-note/actions/workflows/build.yml)

A tiny markdown notetaking system centered around being as lightweight as
possible.

## Concepts

- Note: a singular note within the program that maps to a file
- Private: a separate set of notes that are hidden from the default view (see
  note below)

> [!IMPORTANT]
> For private notes, there is a per-note password enforcement for encrypting
> them, but it should not be relied on for anything beyond basic protection
> against utilities like `cat` or `less`; if you need true encryption security,
> use a trusted first- or third-party solution (i.e., FileVault or VeraCrypt).

## Installation

> [!NOTE]
> This program currently only supports MacOS and Linux distros. It has not been
> thoroughly tested on Linux systems yet.

```sh
curl -sSL https://raw.githubusercontent.com/lxsavage/pico-note/refs/heads/main/scripts/install.sh | bash
```

## Uninstall

```sh
curl -sSL https://raw.githubusercontent.com/lxsavage/pico-note/refs/heads/main/scripts/uninstall.sh | bash
```

### Manual installation

In order to build this project, the Golang CLI needs to be installed and on
path. For more information on how to do this, check the
[Golang install guide](https://go.dev/doc/install).

To build and install, use `make install`.

> [!NOTE]
> By default, this will be installed under `/usr/local/bin`. This can be changed
> by adjusting the makefile `INSTALL_DIR` variable to the intended path before
> running any of these make commands.

Upgrading from a previous version is as simple as pulling the latest changes,
then running `make upgrade`.

---

Uninstallation is just `make uninstall`.

## Usage

```sh
piconote [--private] [-h] <command> [<file>]
```

- `--private` sets the program to use the private set of notes instead of the
  standard set
- `-h` show the usage of the program

### Available commands:

- `view` view an existing note; will output nothing if the note doesn't exist
- `list` list all available notes in the current area (private/non-private)
- `write` open a note in the system editor; will create if not existing yet
- `remove` removes an existing note; does nothing if note does not already
  exist
