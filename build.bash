


go get github.com/mhale/smtpd
go install github.com/mhale/smtpd

go get github.com/go-resty/resty   
go install github.com/go-resty/resty 

go get github.com/nsqio/go-nsq
go install github.com/nsqio/go-nsq

go get github.com/go-sql-driver/mysql
go install github.com/go-sql-driver/mysql


go install bitbucket.org/ix-specops/golangeventserver/src/main

mv $GOPATH/bin/main $GOPATH/bin/truevieweventserver 
cp $GOPATH/bin/truevieweventserver $GOPATH/src/bitbucket.org/ix-specops/golangeventserver/dist

echo "build was successful"

