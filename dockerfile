# 基础镜像
FROM ubuntu:jammy
# 备份原始安装源
RUN cp /etc/apt/sources.list /etc/apt/sources.list.bak
# 修改为国内源
RUN sed -i 's/archive.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
RUN apt update
RUN apt full-upgrade
RUN apt install -y translate-shell mediainfo sqlite3 git build-essential bsdmainutils
# 安装基础软件
WORKDIR /root
RUN wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
RUN export PATH=$PATH:/usr/local/go/bin
# go env
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GOBIN=/go/bin
RUN export CGO_ENABLED=1
COPY srt /usr/local/bin/srt
# CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o srt main.go
# 准备软件
CMD ["srt"]