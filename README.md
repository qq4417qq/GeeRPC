这是一个简单的rpc框架
通过
```shell
docker build -t geerpc:v1 .
```
可以对项目快速进行简单的docker构建
```shell
#可以加载镜像
docker load geerpc
```
```shell
#可以运行镜像
docker run -it geerpc:v1
```
```shell
#查看正在运行的镜像
docker ps -a
```
```shell
#停止正在运行的所有镜像
docker stop *
```
```shell
#删除NAME为cc0add06787f的镜像
docker rm cc0add06787f
```
```shell
#将镜像复制一份命名为geerpc
docker save geerpc:v1 > geerpc
```
