SET VERSION=v0.4.3

git tag %VERSION%
git push origin %VERSION%
SET GOPROXY=proxy.golang.org 
go list -m github.com/zeroboo/serverutils@%VERSION%
