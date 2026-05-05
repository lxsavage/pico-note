# PicoNote

A tiny markdown notetaking system centered around being as lightweight as
possible.

## Concepts

- Note: a singular note within the program that maps to a file
- Private: a separate set of notes that are hidden from the default view (see
  note below)

> ![IMPORTANT]
> For private notes, there is a per-note password enforcement for encrypting
> them, but it should not be relied on for anything beyond basic protection
> against utilities like `cat` or `less`; if you need true encryption security,
> use a trusted first- or third-party solution (i.e., FileVault or VeraCrypt).

## Usage

```sh
piconote [-private] [-h] <command> [<file>]
```

- `-private` sets the program to use the private set of notes instead of the
  standard set
- `-h` show the usage of the program

### Available commands:

- `view` view an existing note; will output nothing if the note doesn't exist
- `list` list all available notes in the current area (private/non-private)
- `write` open a note in the system editor; will create if not existing yet
- `remove` removes an existing note; does nothing if note does not already
  exist
