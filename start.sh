#!/bin/bash
number=10
cp /root/.kube/config ./
docker build . -t asinitsyn1024/test9:$number
docker push asinitsyn1024/test9:$number
cat <<EOF > pod.yaml
apiVersion: v1
kind: Pod
metadata:
  name: sleep
spec:
  containers:
  - name: sleep
    image: asinitsyn1024/test9:$number
  imagePullSecrets:
  - name: regcred
EOF
kubectl apply -f pod.yaml
rm config
