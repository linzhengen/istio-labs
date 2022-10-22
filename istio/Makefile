deploy-base:
	helm repo add istio https://istio-release.storage.googleapis.com/charts --force-update
	helm repo update istio
	helm upgrade --install istio-base istio/base \
	--kube-context minikube \
	--create-namespace \
	--namespace istio-system

deploy-istiod:
	helm upgrade --install istiod istio/istiod \
	--kube-context minikube \
	--create-namespace \
	--namespace istio-system
