init:
	minikube addons enable metrics-server
	brew install skaffold
	brew install vegeta

dev:
	skaffold dev -v debug

build:
	skaffold build --file-output output.json

load-test:
	echo 'GET http://localhost:8888/helloWorld' | vegeta attack -rate=30 -duration=120s | vegeta report

