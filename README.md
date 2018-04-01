# go-crawl-store
go-crawl-store a restful API to price crawler for online stores (大潤發, 家樂福)

## Start using it
0. Please note that because of the net/html dependency, goquery requires Go1.1+

1. Download and install it:

```sh
 $ go get github.com/EddieYY/go-crawl-store
```
2. To run (test):

```sh
 $ cd $GOPATH/src/github.com/EiddieYY/go-crawl-store
 $ go run main.go

```
<img align="right" width="159px" src="https://raw.githubusercontent.com/EddieYY/go-crawl-store/master/img/go-crawl-stire_server_run.png">


How to `curl`:

```bash
curl http://127.0.0.1:8080/api/v1/家樂福/醬油
curl http://127.0.0.1:8080/api/v1/大潤發/醬油
curl http://127.0.0.1:8080/api/v1/ALL/醬油
```
