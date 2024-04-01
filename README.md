# go-getting-started

brew search go

brew install go@1.22.1

go version

mkdir go-getting-started
cd go-getting-started

go mod init example/go-getting-started

go get github.com/gin-gonic/gin
go get github.com/joho/godotenv

go run .
lsof -i :4000
kill -9 <pid>

echo $GOPATH
echo $(go env GOPATH)

curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

vi ~/.zshrc
alias air='$(go env GOPATH)/bin/air'
source ~/.zshrc

air init

// to run server with air live-reload:
air
air -c .air.toml 
