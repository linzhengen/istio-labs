deploy:
	kubectl create namespace locust --dry-run=client -o yaml | kubectl apply -f -
	kubectl -n locust create configmap hello-world-loadtest --from-file main.py --dry-run=client -o yaml | kubectl apply -f -
	helm repo add deliveryhero https://charts.deliveryhero.io/ --force-update
	helm repo update deliveryhero
	helm upgrade --install locust deliveryhero/locust \
	--kube-context minikube \
	--create-namespace \
	--namespace locust \
	--set loadtest.name=hello-world-loadtest \
    --set loadtest.locust_locustfile_configmap=hello-world-loadtest
