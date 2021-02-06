## Gerando image do docker

```
docker build -t HRsniper/codepix:latest -f codepix/Dockerfile.prod codepix
```

### Docker hub

```
docker push HRsniper/codepix:latest
```

### Kubernetes

Vamos utilizar [Kind](https://kind.sigs.k8s.io/) é uma ferramenta para executar clusters Kubernetes locais usando "nós" de contêiner do Docker.

```
GO111MODULE="on" go get sigs.k8s.io/kind

kind create cluster --name=codepix
```

### kubectl

A ferramenta de linha de comando do Kubernetes o [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/), permite que você execute comandos em clusters do Kubernetes. Você pode implantar aplicativos, inspecionar e gerenciar recursos de cluster e visualizar registros.

```
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

kubectl cluster-info --context kind-codepix

kubectl get nodes
```

