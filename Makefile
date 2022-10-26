init:
	minikube addons enable metrics-server
	brew install skaffold
	cd istio && make deploy-base
	cd istio && make deploy-istiod
	cd locust && make deploy

dev:
	skaffold dev -v debug
