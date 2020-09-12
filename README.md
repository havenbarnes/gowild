# gowild

#### A simple CLI for recording terminal commands in a shell script

![Go CI](https://github.com/havenbarnes/gowild/workflows/Go%20CI/badge.svg)

## Installing

First, download the appropriate binary at [github.com/havenbarnes/gowild/releases](github.com/havenbarnes/gowild/releases). If you're using macOS, you want the Darwin distribution.
Unizip the package to continue.

Note: If using macOS, you must first open the executable so that it is trusted by the OS:

- In the Finder on your Mac, locate gowild executable in the directory you just downloaded.
- Control-click the app icon, then choose Open from the shortcut menu.
- Click Open and confirm on any popups.

Now, to add gowild globally, run the following commands:

```bash
$ cd <binary_directory> # e.g. cd gowild_0.0.1_Darwin_x86_64
$ cp gowild /usr/local/bin/gowild
```

## Usage

```bash
$ gowild record
Now recording commands... run 'gowild stop' to end recording
$ echo 'Hello, World!'
$ gowild stop
What should the output file be named? [gowild.sh]: helloworld.sh
$ ./helloworld.sh
Hello, World!
```

## FAQ

### What if I use Zsh instead of Bash?

Zsh is also supported! Fish is not yet supported, but planned.

## License

[ISC](LICENSE)
