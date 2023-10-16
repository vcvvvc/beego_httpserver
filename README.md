# beego_httpserver
a blog

---

## https

```editorconfig
//app.conf
EnableDocs = true
EnableHTTPS = true
EnableHttpTLS = true
HTTPSCertFile = "conf/server.crt"
HTTPSKeyFile = "conf/server.key"
```

```
//输入一个简单密码
openssl genrsa -des3 -out server.key 1024 

//密码验证后，common选项输入主机名或server ip
openssl req -new -key server.key -out server.csr

cp server.key server.key.org 

//输入密码验证后生成一份无密码的key
openssl rsa -in server.key.org -out server.key 

//生成自签名的服务器证书
openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.crt
```