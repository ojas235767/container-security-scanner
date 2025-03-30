# Kubernetes Security Scan

## Objective
Scan a local Kubernetes cluster for vulnerabilities and generate a findings report in JSON format.

## Setup Instructions
1. Install Kubernetes cluster using Minikube, K3s, or Kind.
2. Install Kubescape: `curl -s https://raw.githubusercontent.com/kubescape/kubescape/master/install.sh | /bin/bash`
3. Scan and export findings: `kubescape scan --format json --output results.json`
