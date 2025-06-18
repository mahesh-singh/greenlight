#!/bin/bash

source .env



echo "Cleaning up any existing ESO installations..."
kubectl delete -f k8s/eso/secretstore.yaml --ignore-not-found
kubectl delete -f k8s/eso/externalsecret.yaml --ignore-not-found
kubectl delete deployment external-secrets-operator -n security --ignore-not-found
kubectl delete deployment external-secrets -n external-secrets-system --ignore-not-found
kubectl delete namespace external-secrets-system --ignore-not-found
kubectl delete crd -l app.kubernetes.io/name=external-secrets --ignore-not-found


# Create namespaces
echo "Creating namespaces..."
kubectl apply -f k8s/namespaces/namespaces.yaml

# Deploy Vault
echo "Deploying HashiCorp Vault..."
kubectl apply -f k8s/vault/vault.yaml


echo "Waiting for Vault to be ready..."
sleep 20
kubectl -n security wait --for=condition=available --timeout=300s deployment/vault


# Setup Vault with initial secrets (in a real environment, you would use proper authentication)
echo "Setting up Vault with initial secrets...(may already be enabled in dev mode)..."
# Get the pod name
VAULT_POD=$(kubectl -n security get pods -l app=vault -o jsonpath="{.items[0].metadata.name}")


# Set Vault address to use HTTP and authenticate
echo "Enabling KV secrets engine..."
kubectl -n security exec $VAULT_POD -- /bin/sh -c 'export VAULT_ADDR="http://127.0.0.1:8200" && export VAULT_TOKEN="root" && vault secrets enable -path=secret kv-v2 || echo "KV secrets engine already enabled"'

echo "Creating database credentials in Vault..."
kubectl -n security exec $VAULT_POD -- /bin/sh -c 'export VAULT_ADDR="http://127.0.0.1:8200" && export VAULT_TOKEN="root" && vault kv put secret/database/credentials username=greenlight password=pa55word dsn="postgres://greenlight:pa55word@postgres:5432/greenlight?sslmode=disable"'

echo "Verifying secrets were created..."
kubectl -n security exec $VAULT_POD -- /bin/sh -c 'export VAULT_ADDR="http://127.0.0.1:8200" && export VAULT_TOKEN="root" && vault kv get secret/database/credentials'


# Install ESO CRDs first
echo "Installing External Secrets Operator CRDs..."
kubectl apply -f https://raw.githubusercontent.com/external-secrets/external-secrets/v0.9.11/deploy/crds/bundle.yaml

echo "Waiting for CRDs to be established..."
kubectl wait --for condition=established --timeout=60s crd/secretstores.external-secrets.io
kubectl wait --for condition=established --timeout=60s crd/externalsecrets.external-secrets.io


# Deploy ESO
echo "Deploying External Secrets Operator..."
kubectl apply -f k8s/eso/eso.yaml


echo "Waiting for ESO to be ready..."
kubectl -n security wait --for=condition=available --timeout=300s deployment/external-secrets-operator

# Create SecretStore and ExternalSecret
echo "Creating SecretStore and ExternalSecret..."
kubectl apply -f k8s/eso/secretstore.yaml
kubectl apply -f k8s/eso/externalsecret.yaml

# Wait for the secret to be created
echo "Waiting for the secret to be created..."
kubectl wait --for=condition=Ready --timeout=120s externalsecret/db-credentials -n greenlight-api



# Create secret for database credentials
#kubectl create secret generic db-credentials \
#  --from-literal=postgres-user="${POSTGRES_USER}"  \
#  --from-literal=postgres-password="${POSTGRES_PASSWORD}" \
#  --from-literal=db-dsn="${DB_DSN}"

# Deploy database to Node B (database node)
kubectl apply -f k8s/database/deployment.yaml

# Wait for database to be ready
echo "Waiting for database to be ready..."
kubectl wait --for=condition=available --timeout=300s deployment/postgres -n greenlight-api

# Create configmap for migrations
kubectl create configmap migrations-configmap --from-file=migrations/ -n greenlight-api


# Run database migrations
echo "Creating database migrations ConfigMap..."
kubectl apply -f k8s/database/migration.yaml

echo "Running database migrations..."
kubectl -n greenlight-api wait --for=condition=complete --timeout=300s job/db-migration 

# Deploy the API application
echo "Deploying Greenlight API..."
kubectl apply -f k8s/application/deployment.yaml

echo "Waiting for Greenlight API to be ready..."
kubectl -n greenlight-api wait --for=condition=available --timeout=300s deployment/greenlight-api

# Deploy the Ingress
echo "Deploying Ingress..."
kubectl apply -f k8s/application/ingress.yaml

# Deploy observability stack
echo "Deploying observability stack..."
kubectl apply -f k8s/dependent-services/observability.yaml

echo "Waiting for observability stack to be ready..."
kubectl -n observability wait --for=condition=available --timeout=300s deployment/prometheus
kubectl -n observability wait --for=condition=available --timeout=300s deployment/grafana

# Setup port forwarding for easy access
echo "Setting up port forwarding..."
kubectl -n greenlight-api port-forward svc/greenlight-api 4000:80 &
kubectl -n observability port-forward svc/grafana 3000:3000 &
kubectl -n observability port-forward svc/prometheus 9090:9090 &

echo "Deployment complete!"
echo "API is accessible at http://localhost:4000"
echo "Grafana is accessible at http://localhost:3000"
echo "Prometheus is accessible at http://localhost:9090"
echo ""
echo "To access the API using the Ingress, add the following to your /etc/hosts file:"
echo "127.0.0.1 greenlight.kube"
echo "Then you can access the API at http://greenlight.kube"
echo "------"
echo "if using docker desktop run the sudo minikube tunnel and then API should be available curl -H "Host: greenlight.kube" http://127.0.0.1/v1/healthcheck"