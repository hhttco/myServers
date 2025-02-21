# myServers
这个一个服务器探针程序

# 面板一键安装脚本

    bash <(curl -Ls https://raw.githubusercontent.com/hhttco/shell/refs/heads/main/MyServersTz/my-servers-install.sh)

# 被控端
    bash <(curl -Ls https://raw.githubusercontent.com/hhttco/shell/refs/heads/main/MyServersTz/my-servers-install-client.sh) 参数1 参数2 参数3

参数1 = auth_secret密钥  
参数2 = 服务端地址，例如：wss://后端IP:123.123.123.123或者后端域名/my_tz  
参数3 = 被控服务器的名字 (不可重复,不可以有空格,开头应该以服务器地区开头，比如：HK001 如果不以这个开头无法获取到服务器地区)
