## shasumgo

`checksum` go.

### install

```sh
$ go get github.com/syui/shasumgo
```

### use

```sh
$ shasumgo ~/.bashrc ~/.zshrc
da39a3ee5e6b4b0d3255bfef95601890afd80709 336de4a66576637c35e0a02e3b4233fdb28798db

$ shasumgo ~/.bashrc ~/.bashrc
ok
$ shasumgo s da39a3ee5e6b4b0d3255bfef95601890afd80709 ~/.bashrc
ok

# success:result display is none
$ shasumgo c ~/.bashrc ~/.bashrc
$ shasumgo s c da39a3ee5e6b4b0d3255bfef95601890afd80709 ~/.bashrc
```

### option

```
# version
$ shasumgo -v

# help
$ shasumgo h
```
