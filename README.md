# TransInDocker

Translate-shell in docker

需要一个proxy变量
~~需要一个language变量~~

# usage

```bash

docker run -dit --rm --name=trans -v '/c/Users/zen/Videos/test:/data' -e proxy=192.168.1.20:8889 trans:latest
docker run -dit--name=trans -v '/f/alist/电影:/srt' -e proxy=192.168.1.5:8889 trans:latest srt
docker build --no-cache -t trans:latest .
```