

deploy_app:
	helm package ping-takehome
	helm install ping-takehome ./ping-takehome-0.1.5.tgz --namespace ping-takehome --create-namespace
	sleep 5
	minikube addons enable ingress
	sleep 5
	echo "127.0.0.1 ping-takehome.local" | sudo tee -a /etc/hosts
	minikube tunnel

all: deploy_app