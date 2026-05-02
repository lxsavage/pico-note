# PicoNote

A tiny markdown notetaking system centered around being as lightweight as
possible.

## Concepts

- Note: a singular note within the program that maps to a file
- Private: a separate set of notes that are hidden from the default view (these
  are not encrypted, so don't put confidential info here that shouldn't be seen
  by another user with access to your home directory)

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
