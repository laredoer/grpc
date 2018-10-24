#私钥
openssl ecparam -genkey -name secp384r1 -out server.key
#自签公钥
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650


# Country Name (2 letter code) [国家]:
# State or Province Name (full name) [省份]:
# Locality Name (eg, city) [城市]:
# Organization Name (eg, company) [公司]:
# Organizational Unit Name (eg, section) [部门]:
# Common Name (eg, fully qualified host name) [一般填写你的服务名称]:go-grpc-example #  credentials.NewClientTLSFromFile("../conf/server.pem", "go-grpc-example")会用到
# Email Address [邮箱]: