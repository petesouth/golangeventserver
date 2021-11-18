	# README #

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Basically.. Things I wanted done on a timer..
*
* Version
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###

* Summary of set up

* After setting up go.  Get the go get the following: 
(tested with `pkg install go` on FreeBSD 10.3, 11, and -CURRENT)
NOTE: Make sure GOPATH is set correctly:
export GOPATH=Full Path to GO Src Root

cd $GOPATH

* go get github.com/go-resty/resty   
* go install github.com/go-resty/resty 

* go get github.com/nsqio/go-nsq
* go install github.com/nsqio/go-nsq

go get github.com/go-sql-driver/mysql
go install github.com/go-sql-driver/mysql



* cd $GOPATH
* cd src 
* mkdir -r bitbucket.org/ix-specops
* cd bitbucket.org/ix-specops
* git clone https://bitbucket.org/ix-specops/golangeventserver  
* cd $GOPATH     So your at root again    
* export GOPATH=$GOPATH

* go install bitbucket.org/ix-specops/golangeventserver/src/main 
* cd $GOPATH/bin 

* Type ./main   You'll see the output asking for the location of a json config file.  


* Configuration

see ./trueviewdemo-ixconfig.json


* Dependencies

github.com/go-resty/resty

github.com/nsqio/go-nsq

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###


* Repo owner or admin
* Other community or team contact


https://bitbucket.org/ix-specops/golangeventserver


