# graphql-vue-example

A try out project with Golang as a GraphQL backend and Vue.js as front-end.

## Run golang project
```
go get github.com/graphql-go/graphql
go build
./graphql-vue-example
```

For active development, you might want to use gin for hot reloading instead of `./graphql-vue-example`.
```
https://github.com/codegangsta/gin
gin -a 8000 -i --all
```

## Run front-end

```
cd frontend
npm install    #if you already didn't do this
npm run dev    #development
npm run build  #build production
```

[Blog post](https://jelinden.blogspot.fi/2017/06/trying-out-graphql-with-golang-and-vuejs.html)
