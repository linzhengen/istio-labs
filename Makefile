dev:
	skaffold dev -v debug

build:
	skaffold build --file-output output.json

load-test:
	echo 'GET http://localhost:8888/helloWorld' | vegeta attack -rate=20 -duration=60s >