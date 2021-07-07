# kcc-toolkit
Toolkit for KCC

**Ubuntu:**

Note: Some flavors of Ubuntu have an out of date go package, so you must add the updated repo manually

````
$ sudo add-apt-repository ppa:longsleep/golang-backports
$ sudo apt update
$ sudo apt install golang-go
````

````
$ git clone https://github.com/kucoin-community-chain/kcc-toolkit.git
$ cd kcc-toolkit
$ go get -d -v
$ ./build.sh
````