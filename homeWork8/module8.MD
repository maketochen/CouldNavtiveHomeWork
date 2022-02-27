# 快捷指令
alias k=kubectl

# 创建httpserver-deployment
k create -f httpserver-deploy.yaml

# 创建httpserver-service
k create -f httpserver-service.yaml

# 创建nginx-ingress-deployment
kubectl create -f nginx-ingress-deployment.yaml

# openssl -x509
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=*/O=cncamp"
Generating a RSA private key

# 创建secret
k create -f secret.yaml

# 创建ingress
k create -f ingress.yaml

