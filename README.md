## jwt tool

**this is a simple tool for parse JWT,and this command read from stdin.**

**install jwt form release and mv it to `$PATH` on your system!**

>Tips: when you working with kubernetes serviceAccount,sometimes you can parse the secret token 

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



#### kubernetes Token Parse

```shell
]# kubectl create sa zhangsan
]# token=`kubectl create token zhangsan` && echo $token

example token:
eyJhbGciOiJSUzI1NiIsImtpZCI6IjR1R0hUWTVnZGg1dDZMWXI3dmlzWm9pQ0ZUeXliS2ZXaHFQa0ljUnJXeTQifQ.eyJhdWQiOlsiaHR0cHM6Ly9rdWJlcm5ldGVzLmRlZmF1bHQuc3ZjLmNsdXN0ZXIubG9jYWwiXSwiZXhwIjoxNjgxMzgwNDI1LCJpYXQiOjE2ODEzNzY4MjUsImlzcyI6Imh0dHBzOi8va3ViZXJuZXRlcy5kZWZhdWx0LnN2Yy5jbHVzdGVyLmxvY2FsIiwia3ViZXJuZXRlcy5pbyI6eyJuYW1lc3BhY2UiOiJkZWZhdWx0Iiwic2VydmljZWFjY291bnQiOnsibmFtZSI6InpoYW5nc2FuIiwidWlkIjoiZjA4NTUwZGEtY2Q2YS00YmU5LWJjYTctNWFmYWFiMGM4MTRjIn19LCJuYmYiOjE2ODEzNzY4MjUsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OnpoYW5nc2FuIn0.PJRs31Ec-uOWP2UR6z3Rd2PmVsRyEU__X74rN7mLLNTX8qc_dtiZe_kvDFdZN6FCLlVtdOBox-XYBgD3aLy2I8HKSoEKKOCf7vYzMO7NdOR3kADT86QHTi4DpKyEyUcjydQLb6SP_RIWVwk0H9Nve9VIm22XKEGQplhF9ky0hHdSmdVC1LwqDf9TTASrIPZioj12Rf49H9d5_LumhWk1AmvBBOO9J5MRnU-stKeELnygZb7qpYRE0dbIOJ5o1kvB9OI1pfFtkkf7Lhwo_feylMIjJ_an092xRhQIlpok8W-1NJGUob1lAXTigvWeXEAoxyAUe_PWGAuHIZ6xm6TYfw

]# echo $token | jwt -t 
```

