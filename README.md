# sleepuntil

`sleepuntil` is a simple CLI utility that waits/sleeps until a specified time. It functions like the traditional sleep command but allows users to specify an exact time instead of a duration.

## Installation

To install `sleepuntil`, follow these steps:

1. Clone the repository:

```
git clone https://github.com/kkty/sleepuntil
cd sleepuntil
```

2. Build the application:

```
go build sleepuntil.go
```

This will produce an executable named `sleepuntil` in the current directory.

## Usage

```
./sleepuntil [-v] <time in format H:MM[:SS][am/pm] or HH:MM[:SS] (24h)>
```

- `-v`: Optional. Display how long the program will wait.
- `<time>`: Specify the time you want the program to sleep until.

For instance:

- `./sleepuntil 7:00am` will sleep until 7:00 AM.
- `./sleepuntil -v 19:00` will display how long it will sleep and then sleep until 7:00 PM (19:00).

## Features

- Supports both 12-hour (with am/pm) and 24-hour formats.
- Optionally display the waiting duration with `-v` flag.
- Gracefully handles invalid input formats.
