deploy-prometheus:
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts/ --force-update
	helm repo update prometheus-community
	helm upgrade --install prometheus prometheus-community/prometheus \
	--kube-context minikube \
	--create-namespace \
	--namespace prometheus

