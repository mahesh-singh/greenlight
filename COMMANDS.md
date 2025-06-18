# Kubernetes Commands Reference Guide

## Table of Contents
- [Kubernetes Commands Reference Guide](#kubernetes-commands-reference-guide)
  - [Table of Contents](#table-of-contents)
  - [Cluster Information and Management](#cluster-information-and-management)
  - [Resource Management](#resource-management)
    - [Listing Resources](#listing-resources)
    - [Detailed Resource Information](#detailed-resource-information)
  - [Pod Operations](#pod-operations)
    - [Managing Pods](#managing-pods)
    - [Executing Commands in Containers](#executing-commands-in-containers)
    - [Container Logs](#container-logs)
    - [Creating Temporary Pods for Testing](#creating-temporary-pods-for-testing)
  - [Service Management](#service-management)
  - [Deployment Management](#deployment-management)
  - [Ingress Management](#ingress-management)
  - [Networking Commands](#networking-commands)
    - [Port Forwarding](#port-forwarding)
    - [Minikube Networking](#minikube-networking)
    - [Testing Connectivity](#testing-connectivity)
  - [Configuration Management](#configuration-management)
  - [Namespace Management](#namespace-management)
  - [Resource Monitoring](#resource-monitoring)
  - [Troubleshooting and Debugging](#troubleshooting-and-debugging)
  - [Application Deployment](#application-deployment)
  - [Context Management](#context-management)
  - [Minikube Specific Commands](#minikube-specific-commands)
  - [NGINX Ingress Controller Commands](#nginx-ingress-controller-commands)

## Cluster Information and Management

| Command | Description |
|---------|-------------|
| `kubectl config get-contexts` | List all contexts |
| `kubectl config current-context` | Show the current context |
| `kubectl config use-context <context-name>` | Switch to a different context |

## Resource Management

### Listing Resources

| Command | Description |
|---------|-------------|
| `kubectl get <resource>` | List resources in the current namespace |
| `kubectl get <resource> -n <namespace>` | List resources in a specific namespace |
| `kubectl get <resource> --all-namespaces` | List resources across all namespaces |
| `kubectl get <resource> <name> -o yaml` | Show resource in YAML format |
| `kubectl get <resource> <name> -o jsonpath='{<path>}'` | Extract specific fields using JSONPath |

### Detailed Resource Information

| Command | Description |
|---------|-------------|
| `kubectl describe <resource> <name>` | Show detailed information about a resource |
| `kubectl describe <resource> <name> -n <namespace>` | Show detailed information in a specific namespace |

## Pod Operations

### Managing Pods

| Command | Description |
|---------|-------------|
| `kubectl get pods` | List all pods in the current namespace |
| `kubectl get pods -n <namespace>` | List all pods in a specific namespace |
| `kubectl get pods --all-namespaces` | List all pods in all namespaces |
| `kubectl describe pod <pod-name>` | Show detailed information about a pod |
| `kubectl delete pod <pod-name>` | Delete a specific pod |

### Executing Commands in Containers

| Command | Description |
|---------|-------------|
| `kubectl exec -it <pod-name> -- <command>` | Execute a command in a container |
| `kubectl exec -it <pod-name> -- sh` | Open an interactive shell in a pod |
| `kubectl exec -it -n <namespace> deploy/<deployment-name> -- <command>` | Execute a command in a deployment's container |
| `kubectl exec -it <pod-name> -- cat <file-path>` | View container files |
| `kubectl exec -it <pod-name> -- cat <file-path> \| grep <pattern>` | Filter output of files using grep |

### Container Logs

| Command | Description |
|---------|-------------|
| `kubectl logs <pod-name>` | View logs for a pod |
| `kubectl logs -f <pod-name>` | Stream logs for a pod (follow) |
| `kubectl logs -n <namespace> deploy/<deployment-name>` | View logs from a deployment |
| `kubectl logs <pod-name> --previous` | View logs from a previous instance of a pod |
| `kubectl logs -n <namespace> -l <label-selector>` | View logs for pods matching a label selector |

### Creating Temporary Pods for Testing

| Command | Description |
|---------|-------------|
| `kubectl run debug --image=curlimages/curl -it --rm -- sh` | Create a temporary debug pod with curl |
| `kubectl run test-curl --image=curlimages/curl -it --rm -- sh -c "<command>"` | Create a pod, run a command, and delete when finished |
| `kubectl run debug-pod --image=curlimages/curl --restart=Never -it --rm -- sh` | Run a temporary debug pod with curl and open a shell |
| `kubectl run netshoot --image=nicolaka/netshoot --restart=Never -it --rm -- bash` | Run a temporary network troubleshooting pod |

## Service Management

| Command | Description |
|---------|-------------|
| `kubectl get svc` | List all services in the current namespace |
| `kubectl get svc -n <namespace>` | List all services in a specific namespace |
| `kubectl get svc --all-namespaces` | List services across all namespaces |
| `kubectl describe service <service-name>` | Show detailed information about a service |
| `kubectl expose deployment <deployment-name> --port=<port>` | Create a service to expose a deployment |
| `kubectl get endpoints <service-name>` | Check if a service has endpoints |
| `curl -v http://<service-name>.<namespace>.svc.cluster.local:<port>/` | Access a service within the cluster |

## Deployment Management

| Command | Description |
|---------|-------------|
| `kubectl get deployments` | List all deployments in the current namespace |
| `kubectl get deployments -n <namespace>` | List all deployments in a specific namespace |
| `kubectl describe deployment <deployment-name>` | Show detailed information about a deployment |
| `kubectl create deployment <name> --image=<image>` | Create a deployment with the specified image |
| `kubectl scale deployment <deployment-name> --replicas=<number>` | Scale a deployment to a specific number of replicas |
| `kubectl rollout status deployment/<deployment-name>` | Check the status of a deployment rollout |
| `kubectl rollout history deployment/<deployment-name>` | View the rollout history of a deployment |
| `kubectl rollout restart deployment/<deployment-name>` | Triggers a restart of all pods in the deployment |
| `kubectl rollout undo deployment/<deployment-name>` | Rollback to the previous deployment version |
| `kubectl wait --for=condition=ready pod --selector=<selector> --timeout=<seconds>` | Wait until pods matching the selector are ready |

## Ingress Management

| Command | Description |
|---------|-------------|
| `kubectl get ingress` | List all ingress resources in the current namespace |
| `kubectl get ingress -n <namespace>` | List all ingress resources in a specific namespace |
| `kubectl get ingress --all-namespaces` | List ingress across all namespaces |
| `kubectl describe ingress <ingress-name>` | Show detailed information about an ingress |
| `kubectl edit ingress <ingress-name>` | Edit an ingress resource |
| `kubectl annotate ingress <ingress-name> <annotation-key>="<value>" --overwrite` | Add or update an annotation on an ingress |

## Networking Commands

### Port Forwarding

| Command | Description |
|---------|-------------|
| `kubectl port-forward svc/<service-name> <local-port>:<service-port>` | Forward a local port to a service |
| `kubectl port-forward <pod-name> <local-port>:<pod-port>` | Forward a local port to a specific pod |

### Minikube Networking

| Command | Description |
|---------|-------------|
| `sudo minikube tunnel` | Create a network tunnel for LoadBalancer and Ingress services |
| `export NODE_PORT=$(kubectl get -n <namespace> service/<service-name> -o jsonpath='{.spec.ports[0].nodePort}')` | Get NodePort of a service |

### Testing Connectivity

| Command | Description |
|---------|-------------|
| `curl -v http://<host>/path` | Verbose output of curl request to test connectivity |
| `curl -v --http1.1 http://<host>/path` | Force HTTP/1.1 for the request |
| `curl -v -H "Host: <host>" http://<ip>:<port>/path` | Set a custom Host header for the request |
| `curl -v --connect-timeout <seconds> --max-time <seconds> http://<host>/path` | Set connection and total timeout values |

## Configuration Management

| Command | Description |
|---------|-------------|
| `kubectl apply -f <yaml-file>` | Apply a configuration from a file |
| `kubectl get configmaps` | List all ConfigMaps in the current namespace |
| `kubectl get secrets` | List all Secrets in the current namespace |
| `kubectl describe configmap <configmap-name>` | Show detailed information about a ConfigMap |
| `kubectl describe secret <secret-name>` | Show detailed information about a Secret |
| `kubectl patch <resource> <name> --patch "$(cat <patch-file>)"` | Apply a patch to modify a resource |

## Namespace Management

| Command | Description |
|---------|-------------|
| `kubectl get namespaces` | List all namespaces |
| `kubectl create namespace <namespace-name>` | Create a new namespace |
| `kubectl delete namespace <namespace-name>` | Delete a namespace |
| `kubectl config set-context --current --namespace=<namespace-name>` | Switch to a different namespace |

## Resource Monitoring

| Command | Description |
|---------|-------------|
| `kubectl top nodes` | Show CPU/Memory usage of nodes |
| `kubectl top pods` | Show CPU/Memory usage of pods |
| `kubectl get events` | List recent events in the current namespace |

## Troubleshooting and Debugging

| Command | Description |
|---------|-------------|
| `kubectl get pod <pod-name> -o jsonpath='{.status.podIP}'` | Get the IP address of a pod |
| `kubectl get pod <pod-name> -o jsonpath='{.spec.containers[*].ports[*].containerPort}'` | Get the container ports of a pod |
| `kubectl get pods -l <label-selector>` | Get pods matching a label selector |
| `POD_NAME=$(kubectl get pods -l <label-selector> -o jsonpath='{.items[0].metadata.name}')` | Get the first pod name matching a label selector |
| `netstat -tulpn \| grep <port>` | Check if a port is open and what process is using it |

## Application Deployment

| Command | Description |
|---------|-------------|
| `kubectl apply -f <yaml-file>` | Create or update resources defined in a YAML file |
| `kubectl delete -f <yaml-file>` | Delete resources defined in a YAML file |
| `kubectl create -f <yaml-file>` | Create resources defined in a YAML file |
| `kubectl replace -f <yaml-file>` | Replace resources defined in a YAML file |

## Context Management

| Command | Description |
|---------|-------------|
| `kubectl config set-context <context-name> --namespace=<namespace-name>` | Set the default namespace for a context |

## Minikube Specific Commands

| Command | Description |
|---------|-------------|
| `minikube addons list` | List all available Minikube addons |
| `minikube addons enable <addon-name>` | Enable a Minikube addon (e.g., ingress) |
| `minikube addons disable <addon-name>` | Disable a Minikube addon |

## NGINX Ingress Controller Commands

| Command | Description |
|---------|-------------|
| `kubectl get pods -n ingress-nginx` | List all pods in the ingress-nginx namespace |
| `kubectl logs -n ingress-nginx -l app.kubernetes.io/name=ingress-nginx` | View logs for the NGINX Ingress Controller |
| `kubectl get svc -n ingress-nginx` | List all services in the ingress-nginx namespace |