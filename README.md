![Screenshot](resources/shelf_icon.png)
# Shelf
Shelf is a small CLI application that automates those tasks every developer does at the start of the day.
## It supports:
- Opening an IDE/editor by giving its binary location and optionally a particular workspace/folder.
- Pulling concurrently a list of comma separated git repositories via goroutines.
- Starting a server by setting its path and the command to execute.

To use it first add a project configuration via `shelf add` and follow the instructions.
Afterwards the previously added project can be loaded by using `shelf pick {project_name}`.

## Disclaimer:
This has only been tested with Eclipse as IDE and currently only supports Mac OSX as it relies on the `open` command which is not present in Windows/Linux.
MultiOS support will be added in the future.