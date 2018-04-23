[![Build Status](https://travis-ci.org/blockcdn-go/home.svg?branch=master)](https://travis-ci.org/blockcdn-go/home)

# home
Go library for detecting and expanding the user's home directory. 

参考了[go-homedir](https://github.com/mitchellh/go-homedir)库

***为什么不使用os/user**？因为内置的`os/user`包在Darwin系统下需要使用cgo，因此无法进行交叉编译。
