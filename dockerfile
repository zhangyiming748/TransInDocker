# 基础镜像
FROM ubuntu:jammy
# 备份原始安装源
RUN cp /etc/apt/sources.list /etc/apt/sources.list.bak
# 修改为国内源
# RUN sed -i 's/archive.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
# RUN sed -i 's/security.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list
RUN apt update
RUN DEBIAN_FRONTEND=noninteractive apt install -y vim nano less ca-certificates man-db git wget curl language-pack-zh-hans net-tools apt-utils iproute2 dialog file xz-utils vim nano less ca-certificates  locales tzdata iproute2 translate-shell mediainfo sqlite3 git build-essential bsdmainutils
RUN DEBIAN_FRONTEND=noninteractive apt full-upgrade -y
RUN locale-gen zh_CN.utf8
RUN echo "export LC_ALL=zh_CN.UTF-8">> /etc/profile
COPY srt /usr/local/bin
# 安装基础软件
# RUN wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz
# RUN tar -C /usr/local -xzf go1.22.1.linux-amd64.tar.gz
# RUN export PATH=$PATH:/usr/local/go/bin
# # go env
# RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
# RUN go env -w GOBIN=/go/bin
# RUN export CGO_ENABLED=1
# CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o srt main.go
# 准备软件
CMD ["srt"]