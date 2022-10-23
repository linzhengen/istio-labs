init:
	minikube addons enable metrics-server
	brew install skaffold
	brew install vegeta
	cd istio && make deploy-base
	cd istio && make deploy-istiod
	cd locust && make deploy

dev:
	skaffold dev -v debug

load-test:
	kubectl --namespace locust port-forward service/locust 8089:8089


