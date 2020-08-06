# qrgen

`qrgen` lets you generate a QR code in the terminal, from stdin.

Why?

I use a file-sharing service ([share.schollz.com](https://share.schollz.com)) I wrote called [share](https://github.com/schollz/share). It lets you share files directly from the command-line. I wanted a way to generate a QR in the terminal so that I can just take a picture with my phone to go to the resulting link.

Now I can share directly from my terminal to my phone easily.

## Usage

First install with

```
> go get github.com/schollz/qrgen
```

Then, I have an alias in my `~/.zshrc` for sharing files to https://share.schollz.com:

```
alias share='f() { curl --progress-bar --upload-file "$1" https://share.schollz.com | qrgen };f'
```


![Example](https://user-images.githubusercontent.com/6550035/89547118-387fa600-d7ba-11ea-8dc1-0e689c5c20ff.png)


## Acknowledgements

Code here was forked from github.com/mdp/qrterminal, and adopted for stdin (under MIT, Copyright 2019 Mark Percival <m@mdp.im>).

## License

MIT
