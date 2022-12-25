#!/bin/bash
number=1
cp /root/.kube/config ./
docker build . -t asinitsyn1024/test9:$number
docker push asinitsyn1024/test9:$number
cat > pod.yaml << "EOF"
apiVersion: v1
kind: Pod
metadata:
  name: sleep
spec:
  containers:
  - name: sleep
    image: asinitsyn1024/test9:1
  imagePullSecrets:
  - name: regcred
EOF
kubectl apply -f pod.yaml