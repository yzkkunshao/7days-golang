touch 命令
touch命令参数可更改文档或目录的日期时间，包括存取时间和更改时间。
命令格式
touch [选项] 文件
常用参数

-a 或--time=atime或--time=access或--time=use  只更改存取时间
-c 或--no-create  不建立任何文档
-d  使用指定的日期时间，而非现在的时间
-f  此参数将忽略不予处理，仅负责解决BSD版本touch指令的兼容性问题
-m 或--time=mtime或--time=modify  只更改变动时间
-r  把指定文档或目录的日期时间，统统设成和参考文档或目录的日期时间相同 -t  使用指定的日期时间，而非现在的时间

使用实例
1.创建不存在的文件test.txt
 代码解读复制代码touch test.txt
2.更新 test.txt 的实践和 test1.txt 时间戳相同
 代码解读复制代码touch -r test.txt test1.txt

检查 wget 是否已安装
在 Zsh 中输入以下命令检查 wget 是否已经安装：

Bash
which wget
请谨慎使用代码。

如果已经安装，会显示 wget 的路径，如果没有，则需要进行安装。

2. 安装 wget
Ubuntu/Debian 系统：
Bash
sudo apt install wget


linux ~是 /usr/用户名/


top 看cpu和内存占用

