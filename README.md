# GlanHx是什么

最初的目的是想做一个功能比较全的HOST碰撞工具，但在代码中我只将host碰撞作为一个模块去调用，后续还会更新其他模块。


# Host碰撞模块相关用法
![image](https://user-images.githubusercontent.com/58256808/223015131-c7da2b94-c4b6-45cf-80d6-ceeb1108bb09.png)


./GlanHx hostcolide -h

```text
  -D    Debug Mod if Open output All info //调试模块，在host碰撞中我做了一些hash校验用来排除没有影响的数据，如果使用者想查看全部数据可以添加该参数启动debug
  -F string //一个hash生成相关的因素，目前只将title和statuscode作为目标站点特征，host碰撞出来的数据如果与目标站点特征不匹配，则会将数据进行输出
        Generate hash based on factors to filter junk data,Default title,status_code. Supported title,status_code,length (default "title,status_code")
  -H string //碰撞域名
        Host
  -HF string //从文件中导入碰撞域名
        Host in file
  -I string //碰撞IP
        Nginx Ip
  -IF string //从文件中导入碰撞IP
        Nginx Ip in file
  -O string //将结果保存在指定路径，默认为result.txt
        output  (default "result.txt")
  -P string //添加了端口扫描功能，可以指定端口进行扫描，默认为80,443
        Port default 80,443 (default "80,443")
  -T int //指定爆破的线程数
        Thread default 10 (default 10)
```


./GlanHx hostcolide -I ip -H host

./GlanHx hostcolide -IF ip.txt -HF host.txt

./GlanHx hostcolide -IF ip.txt -HF host.txt -T 3

./GlanHx hostcolide -IF ip.txt -HF host.txt -T3 -P 80,443,8080,8443

./GlanHx hostcolide -IF ip.txt -HF host.txt -T3 -P 80,443,8080,8443 -O result.txt




欢迎提出更多模块的建议和代码中的bug，本项目会长期更新
