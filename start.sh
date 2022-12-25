#!/bin/bash
number=3
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
    image: asinitsyn1024/test9:3
  imagePullSecrets:
  - name: regcred
EOF
kubectl apply -f pod.yaml
rm config