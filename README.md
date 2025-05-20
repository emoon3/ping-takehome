# ping-takehome

Assumptions are that you have a working minikube config (though any kubernetes cluster should work) with kubectl installed an configured to access the minikube cluster. Helm also needs to be installed.

1. Clone the repo
2. Navigate to the top level directory (ping-takehome)
3. Update the ping-takehome/values.yml file replacing the placehold for the key parameter with a working apikey
4. Run chmod+x makefile
5. Run ./makefile
6. After the makefile runs, you should be able to connect to ping-takehome.local/query and get a JSON payload

(If you are not deploying on minikube, run the commands in the makefile)

I thought about setting up vault to store the api key, but that would add more time. I also thought about setting up prometheus and grafana for monitoring or even ArgoCD, but didn't want this to take too long.

I look forward to feedback.

Thanks!
