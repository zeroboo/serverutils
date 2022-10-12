# go-serverutils
Common tools to manage backend service &amp; users
Current release: v0.3.1
## 

1. Build
2. Test

```console
go test -timeout 300s github.com/zeroboo/serverutils -v
```

3. Publish:
    Windows: eg, publish v0.3.3

```console
SET GOPROXY=proxy.golang.org 
go list -m github.com/zeroboo/serverutils@v0.3.5
```