FROM golang:1.21 as builder

WORKDIR /SgBlog
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
#    && go build -o adminServer ./app/admin/cmd/main.go
    && go build -o blogServer ./app/blog/cmd/main.go

FROM alpine:latest

LABEL MAINTAINER="2493381254@qq.com"

WORKDIR /SgBlog

COPY --from=0 /SgBlog ./
#COPY --from=0 /SgBlog/app/admin/cmd/etc/admin.docker.yaml ./
COPY --from=0 /SgBlog/app/blog/cmd/etc/blog.docker.yaml ./


#EXPOSE 8989
#ENTRYPOINT ./adminServer -c admin.docker.yaml
# docker build -t sgblog_admin .

# docker run  -d --name sgblog_admin --network sgblog_network --link mysql-server:mysql --link sg-blog-redis:redis -p 8989:8989 sgblog_admin


EXPOSE 7777
ENTRYPOINT ./blogServer -c blog.docker.yaml

# docker build -t sgblog_blog .

# docker run  -d --name sgblog_blog --network sgblog_network --link mysql-server:mysql --link sg-blog-redis:redis -p 7777:7777 sgblog_blog

# docker network create sgblog_network