# go-serverutils
Common tools to manage backend service &amp; users
Current release: v0.4.1
## 

1. Build
2. Test

```console
go test -timeout 300s github.com/zeroboo/serverutils -v
```

3. Publish:
Example with VERSION=v0.4.1
  - Tag on git

```console
git tag $VERSION
git push origin $VERSION
```
  - Publish go 
```console
SET GOPROXY=proxy.golang.org 
go list -m github.com/zeroboo/serverutils@$VERSION
```