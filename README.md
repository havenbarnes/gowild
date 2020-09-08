# gowild

#### A simple CLI for recording terminal commands in a shell script

![Go CI](https://github.com/havenbarnes/gowild/workflows/Go%20CI/badge.svg)

## Installing

```bash
$ git clone https://github.com/havenbarnes/gowild
$ cd gowild
$ ./install.sh
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
