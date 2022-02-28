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

##TODO
官方的yaml中使用的镜像不能下载，所以要改成dockerhub上的镜像
image: k8s.gcr.io/ingress-nginx/controller:v1.0.0  -> https://hub.docker.com/r/willdockerhub/ingress-nginx-controller
docker pull willdockerhub/ingress-nginx-controller:v1.0.0
    
image: k8s.gcr.io/ingress-nginx/kube-webhook-certgen:v1.0 -> https://hub.docker.com/r/jettech/kube-webhook-certgen/tags
docker pull jettech/kube-webhook-certgen:v1.0.0  

！！！！拉不下来的image 统统替换国内的image 
