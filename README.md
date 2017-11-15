### Hey look, a webapp written in go!

To run locally, assuming you have go installed and are in your $GOPATH:
1. $ go get github.com/samharden/go-web-app
2. $ cd $GOPATH/src/github.com/samharden/go-web-app
3. $ go build main.go
4. $ ./main
5. [open localhost:8080 in your browser]

To deploy to Heroku, you're going to have to do some other stuff:

1. read [this](https://devcenter.heroku.com/articles/go-dependencies-via-govendor)
2. then read [this](https://devcenter.heroku.com/articles/getting-started-with-go#prepare-the-app)
3. in the Procfile, change 'testapp' to 'go-web-app'
4. I use git to deploy, buy hey, you do you.
