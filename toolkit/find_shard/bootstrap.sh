#!/bin/sh

make clean

#安装godep
#go get github.com/tools/godep

echo "downloading dependcies, it may take a few minutes..."
# Test godep install
godep path > /dev/null 2>&1
if [ "$?" = 0 ]; then
    GOPATH=`godep path`:$GOPATH
    godep restore
    make || exit $?
    exit 0
fi

go get -u github.com/c4pt0r/cfg
go get -u github.com/go-sql-driver/mysql

make || exit $?
