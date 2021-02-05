FROM golang

RUN mkdir -p /go/src/github.com/alegd/todo-demo

ADD . /go/src/github.com/alegd/todo-demo

RUN cd /go/src/github.com/alegd/todo-demo && go mod vendor

ENTRYPOINT  watcher -run github.com/alegd/todo-demo/cmd -watch github.com/alegd/todo-demo
