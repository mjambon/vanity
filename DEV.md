Getting started with Go
==

The following are some commands I used to set up my development
environment. If this is confusing, it's probably better to get started
at https://golang.org/doc/install.

Go mode for emacs
--

```
$ sudo apt install golang-mode
```

Set tab width to 2 or whatever you like in your `~/.emacs`:

```
; Tab width for Go
(add-hook 'go-mode-hook
          (lambda ()
            (setq indent-tabs-mode 1)
            (setq tab-width 2)))
```

Go compiler and standard library
--

```
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt update
$ sudo apt install golang-go
```

Yaml library
--

```
$ go get gopkg.in/yaml.v2
```

Working with Yaml
--

Yamllint can be used to validate the syntax a Yaml file (won't check
whether the parsed document conforms to a specific structure).
```
$ sudo apt install yamllint
```

Yq can be used to extract specific elements from a Yaml file. It can
also convert a file to Json, which is useful if you're not familiar
with the Yaml syntax. The following instructions were taken from
https://github.com/mikefarah/yq:
```
sudo add-apt-repository ppa:rmescandon/yq
sudo apt update
sudo apt install yq -y
```

Here's how you can convert Yaml to Json and pretty-print it (requires `jq`):
```
$ yq r -j example.yaml | jq
```