Let’s generate keys & certificates.
Replace demo.mlopshub.com with your domain name or IP address.

```shell
openssl req -x509 \
		-sha256 -days 356 \
		-nodes \
		-newkey rsa:2048 \
		-subj "/CN=dongnguyen.dev/C=US/L=San Fransisco" \
		-keyout rootCA.key -out rootCA.crt
```

We will use the rootCA.keyand rootCA.crt to sign the SSL certificate.

Create Self-Signed Certificates using OpenSSL
Follow the steps given below to create the self-signed certificates. We will sign out certificates using our own root CA created in the previous step.

1. Create the Server Private Key

```shell
openssl genrsa -out server.key 2048
```

2. Create Certificate Signing Request Configuration
   We will create a csr.conf file to have all the information to generate the CSR. Replace dongnguyen.dev with your domain name or IP address.

```shell
cat > csr.conf <<EOF
[ req ]
default_bits = 2048
prompt = no
default_md = sha256
req_extensions = req_ext
distinguished_name = dn

[ dn ]
C = US
ST = California
L = San Fransisco
O = dongnguyen
OU = Dong nguyen
CN = dongnguyen.dev

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = dongnguyen.dev
DNS.2 = *.dongnguyen.dev
EOF
```

3. Generate Certificate Signing Request (CSR) Using Server Private Key
   Now we will generate server.csr using the following command.

```shell
openssl req -new -key server.key -out server.csr -config csr.conf
```

Now our folder should have three files. csr.conf, server.csr and server.key

4. Create a external file
   Execute the following to create cert.conf for the SSL certificate. Replace dongnguyen.dev with your domain name or IP address.

```shell
cat > cert.conf <<EOF

authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
subjectAltName = @alt_names

[alt_names]
DNS.1 = dongnguyen.dev
DNS.2 = *.dongnguyen.dev

EOF
```

5. Generate SSL certificate With self signed CA
   Now, execute the following command to generate the SSL certificate that is signed by the rootCA.crt and rootCA.key created as part of our own Certificate Authority.

```shell
openssl x509 -req \
 -in server.csr \
 -CA rootCA.crt -CAkey rootCA.key \
 -CAcreateserial -out server.crt \
 -days 365 \
 -sha256 -extfile cert.conf
```

The above command will generate server.crt that will be used with our server.key to enable SSL in applications.

6. Install Certificate Authority In Your Browser/OS

You need to install the rootCA.crt in your browser or operating system to avoid the security message that shows up in the browser when using self-signed certificates.

Installing self-signed CA certificates differs in Operating systems. For example, in MAC, you can add the certificate by double-clicking it and adding it to the keychain. Check the respective Operating system guide on installing the certificate.

On linux

```shell
sudo cp rootCA.crt /usr/local/share/ca-certificates/DongnguyenSelfRootCA.crt
sudo update-ca-certificates
```

To remove

```shell
sudo rm /usr/local/share/ca-certificates/DongnguyenSelfRootCA.crt
sudo update-ca-certificates --fresh
```
