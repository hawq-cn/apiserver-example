# Apiserver-example
An apiserver-example created by kube apiserver-builder (https://github.com/kubernetes-incubator/apiserver-builder)

## Run on MacOS
```bash
# download repo in you $GOPATH/src/github.com/hawq-cn/apiserver-example
cd ~/go/src/github.com/hawq-cn/apiserver-example

# set go compile flag
export GOOS=darwin
export GOARCH=amd64

# run (make sure etcd in $PATH)
./bin/apiserver-boot run local --generate=false

# verify
kubectl --kubeconfig kubeconfig api-versions
kubectl --kubeconfig kubeconfig create -f sample/myresource.yaml
kubectl --kubeconfig kubeconfig describe MyResources myresource-example
```

## Key code
1. type: pkg/apis/core/v1alpha1
2. controller: pkg/controller/myresource

## Hot to create it?
See the commits history:
https://github.com/hawq-cn/apiserver-example/commits/master

