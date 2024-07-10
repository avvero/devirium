# SHA256 with ECDSA

To generate private key do
```bash
openssl ecparam -genkey -name prime256v1 -out ec_private.pem
openssl pkcs8 -topk8 -nocrypt -in ec_private.pem -out ec_private_pkcs8.pem
```
Use key from `ec_private_pkcs8.pem` as private key

To generate public key do
```bash
openssl ec -in ec_private.pem -pubout -out ec_public.pem
```

Java обычно предпочитает ключи в формате PKCS#8. 

#sha #cryptography #java
#draft