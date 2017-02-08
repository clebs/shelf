# Shelf
Shelf is a small CLI application that I did for myself to bootup my projects automatically when working.
## It supports:
- Opening an IDE/editor by giving its binary location and optionally a particular workspace/folder.
- Pulling concurrently a list of comma separated git repositories via goroutines.
- Starting a server by setting its path and the command to execute.

## Disclaimer:
This has only been tested with Eclipse as IDE and currently only supports Mac OSX as it relies on the `open` command which is not present in Windows/Linux.
MultiOS support will be added in the future.