deploy-grafana:
	helm repo add grafana https://grafana.github.io/helm-charts --force-update
	helm repo update grafana
	helm upgrade --install grafana grafana/grafana \
	--kube-context minikube \
	--create-namespace \
	--namespace grafana

admin-password:
	kubectl get secret --namespace grafana grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

port-forward:
	kubectl --context minikube -n grafana port-forward service/grafana 3000:80
