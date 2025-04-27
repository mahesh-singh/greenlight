#!/bin/bash

source .env

# Create secret for database credentials
kubectl create secret generic db-credentials \
  --from-literal=postgres-user="${POSTGRES_USER}"  \
  --from-literal=postgres-password="${POSTGRES_PASSWORD}" \
  --from-literal=db-dsn="${DB_DSN}"

# Deploy database to Node B (database node)
kubectl apply -f k8s/database/deployment.yaml

# Wait for database to be ready
echo "Waiting for database to be ready..."
kubectl wait --for=condition=available --timeout=300s deployment/postgres

# Create configmap for migrations
kubectl create configmap migrations-configmap --from-file=migrations/

# Run migrations
kubectl apply -f k8s/database/db-migration-job.yaml

# Deploy API to Node A (application node)
kubectl apply -f k8s/application/deployment.yaml

# Deploy observability stack to Node C (dependent services node)
kubectl create configmap prometheus-config --from-file=k8s/dependent-services/configmap.yaml
kubectl apply -f k8s/dependent-services/observability-stack.yaml

# Check deployment status
echo "Checking deployment status..."
kubectl get pods -o wide

# Set up port forwarding for API access
echo "Setting up port forwarding for API access..."
kubectl port-forward svc/greenlight-api 4000:4000 &

# Set up port forwarding for Grafana
echo "Setting up port forwarding for Grafana access..."
kubectl port-forward svc/grafana 3000:3000 &

echo "Deployment complete!"
echo "API is accessible at http://localhost:4000"
echo "Grafana is accessible at http://localhost:3000"