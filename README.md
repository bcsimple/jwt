## jwt tool

**this is a simple tool for parse JWT,and this command read from stdin,**

#### Options

- -p  指定解析的段落 默认为part1和part2
- -t  人类可读时间
- -v token是否校验
- -s 指明校验secret
- --pub 指明校验公钥文件

#### Usage examples

```shell
1.解析全部
echo $TOKEN | jwt 

2.解析part1并且时间可读
echo $TOKEN | jwt -p -t 

3.解析part2并且时间可读
echo $TOKEN | jwt -p -t 

4.指定secret校验token
echo $TOKEN | jwt -v -s "123"

5.指定pub文件校验token
echo $TOKEN | jwt -v --pub sa.pub
```



