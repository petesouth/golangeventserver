#!/usr/bin/env bash

if [ $# -ne 1 ]; 
then 
    echo "illegal number of parameters.  this utility takes one parameter"
    echo "./setupgo.bash /rootpath/of/goinstall" 
    exit 0
fi


INSTALL_ROOT=$1
export INSTALL_ROOT

echo "runing the go install.  Using $INSTALL_ROOT as the target directory"

pause

GOOS=freebsd
export GOOS

CGO_ENABLED=0
export CGO_ENABLED

GOARCH=amd64
export GOARCH

GOROOT_BOOTSTRAP=$INSTALL_ROOT/go1.4
export GOROOT_BOOTSTRAP

GOROOT=$INSTALL_ROOT/gosrc/go
export GOROOT

GOPATH=$INSTALL_ROOT/go
export GOPATH

#Adding GO BIN to my path.
PATH=./:$PATH:$GOROOT/bin:$GOPATH/bin
export path

run_program=true
export run_program

if [ -d "$GOROOT_BOOTSTRAP" ]
then
	echo "Warning you must move or delete your current $GOROOT_BOOTSTRAP for this installation to continue"
	run_program=false
fi

if [ -d "$GOROOT" ]
then
     echo "Warning you must move or delete your current $GOROOT directory for this installation to continue"
     run_program=false
fi

if [ -d "$GOPATH" ]
then
     echo "Warning you must move or delete your current $GOPATH directory for this installation to continue"
     run_program=false
fi


function printEnv {
  echo "PATH=$PATH"
  echo "GOPATH=$GOPATH"
  echo "GOROOT=$GOROOT"
  echo "GOARCH=$GOARCH"
  echo "GOOS=$GOOS"

  echo "CGO_ENABLED=$CGO_ENABLED"
  echo GOROOT_BOOTSTRAP=$GOROOT_BOOTSTRAP

}

if [ "$run_program" = false ]
then
   echo exiting program.  Pleases make specified modifications.
   printEnv
   exit 1
fi

echo "creating scratch directory (If one is created already it'll be deleted and re-created)"

if [ -d "./scratch" ]
then
    echo "Directory ./scratch exists, going to remove it for recreation"
    rm -rf ./scratch
fi

mkdir ./scratch
cd scratch

echo "currecnt working directory is:"
pwd

echo "setting up go."

echo "Getting go1.4.3 as a bootstrap"

wget https://storage.googleapis.com/golang/go1.4.3.freebsd-amd64.tar.gz
tar -xvf go1.4.3.freebsd-amd64.tar.gz

echo "moving go to $GOROOT_BOOTSTRAP"


mv ./go $GOROOT_BOOTSTRAP

cd ..
pwd

echo "********************************************************"
echo "***> Success: all the go directories I installed are in:" 
echo "********************************************************"

find $INSTALL_ROOT -maxdepth 1 -name "go*"


echo "Now I have go1.4 in place for bootstrapping building a branch of go1.8.3 for freebsd"

pushd .
mkdir -p $GOROOT
cd $GOROOT

git clone https://go.googlesource.com/go .
cd go
git checkout go1.8.3
cd src
./all.bash
./make.bash

popd

echo "**************************************************"
echo "***> Success: Just setup $GOROOT"
echo "***************************************************"

if [ -d "$GOPATH" ]
then
     echo "Your $GOPATH exists.. using it.  If you want a fresh $GOPATH move the old one out of the way"
else
    echo "NO $GOPATH exists creating one"
    mkdir $GOPATH
fi


pushd .
cd $GOPATH

GOPATH=$GOPATH GOROOT=$GOROOT GOOS=$GOOS GOARCH=$GOARCH go get github.com/sparrc/gdm
GOPATH=$GOPATH GOROOT=$GOROOT GOOS=$GOOS GOARCH=$GOARCH go install github.com/sparrc/gdm

GOPATH=$GOPATH GOROOT=$GOROOT GOOS=$GOOS GOARCH=$GOARCH go get github.com/nsf/gocode
GOPATH=$GOPATH GOROOT=$GOROOT GOOS=$GOOS GOARCH=$GOARCH go install github.com/nsf/gocode

popd


echo "********* SUCCESS GO IS INSTALLED ********"
echo "For more information see: https://golang.org/doc/install/source"
echo "The environment is as follows:"
echo "******************************************"

printEnv

echo MAKE SURE YOUR .CSHRC, .BASHRC, .SHRC etc.. are updated for these environment variables. ENJOY !!!

pwd


