# TransInDocker

Translate-shell in docker

需要一个proxy变量
需要一个from变量
需要一个to变量

# usage

```bash
docker run -dit --rm --name=trans -v /path/to/origin/srt:/srt -e proxy=192.168.1.5:8889 -e from=ja -e to=zh trans:latest srt
docker run -dit --rm --name=trans -v /f/large/GirlFriend4ever/G4E:/srt -e proxy=192.168.1.5:8889 -e from=en -e to=zh trans:latest srt
```
