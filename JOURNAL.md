
Next - https://one2n.io/sre-bootcamp/sre-bootcamp-exercises/2-containerise-rest-api

Anki - page - 116


When I executed script the via `ssh -i ~/.ssh/greenlight.pem -t ubuntu@52.90.171.119 "sudo bash /home/ubuntu/setup/01.sh"` looks like it didn't executed the following lines 


`sudo mkdir -p /home/${USERNAME}/.ssh`
`sudo cp -r /home/ubuntu/.ssh/authorized_keys /home/${USERNAME}/.ssh/`
`sudo chown -R ${USERNAME}:${USERNAME} /home/${USERNAME}/.ssh`
`sudo chmod 700 /home/${USERNAME}/.ssh`
`sudo chmod 600 /home/${USERNAME}/.ssh/authorized_keys`

When I ssh as a `ubuntu` user and executed line by line, I am able to do the ssh via new user

SSH hung up after internet discomment, how to end the session 
~.

Htop - for activity monitor 

ls -a for all files including hidden

mac terminal  - ctr a, ctr e

/root/.ssh/authorized_keys

# sed -i -e 's/.*exit 142" \(.*$\)/\1/' /root/.ssh/authorized_keys

/root/.ssh/authorized_keys contains a line `Please login as the user "ubuntu" rather than the user "root".` removing that allow to login via greenlight


https://repost.aws/knowledge-center/new-user-accounts-linux-instance


Copy Setup/01.sh

`rsync -e "ssh -i ~/.ssh/greenlight.pem" -rP --delete ./remote/setup ubuntu@3.93.241.94:/home/ubuntu/`

Execute the script

`ssh -i ~/.ssh/greenlight.pem -t ubuntu@3.93.241.94 "sudo bash /home/ubuntu/setup/01.sh"`

SSH
`ssh -i ~/.ssh/greenlight.pem -t ubuntu@3.93.241.94`





# Migration error
Getting error while running the migration 
- login as a greenlight and run the migration command directly  - `migrate -path ~/migrations -database $GREENLIGHT_DB_DSN up`

Log in as Postgres User
sudo -u postgres psql

Get the DB owner name 
`SELECT datname, pg_catalog.pg_get_userbyid(datdba) AS owner FROM pg_database WHERE datname = 'greenlight';`

Change the ownership of DB
`ALTER DATABASE greenlight OWNER TO greenlight;`

psql meta command 
- `\dt` table name
- `\l` database list


systemd - 
unit file - https://www.digitalocean.com/community/tutorials/understanding-systemd-units-and-unit-files
https://gist.github.com/ageis/f5595e59b1cddb1513d1b425a323db04
journalctl : https://www.digitalocean.com/community/tutorials/how-to-use-journalctl-to-view-and-manipulate-systemd-logs



Docker 
https://claude.ai/chat/38433b2c-8681-42a3-8a71-4742685898ac
https://chatgpt.com/c/67edf0c4-8a24-8008-80a0-951c083962ed
https://claude.ai/chat/7bc41cf4-0641-44c6-92f2-26630a5492bf

docker stats
docker exec
docker history <image>


08/APR/2025: github action  - done, need to test API image


----


 kubectl describe pod greenlight-api-6d9cd9dcd7-2cfvk

 Understand controllers
 kubectl get pods -n ingress-nginx
 kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.1/deploy/static/provider/baremetal/deploy.yaml

minikube profile
`minikube profile greenlight-cluster`
`minikube profile list`

? - curl -H "Host: greenlight.local" http://<Ingress-Controller-IP>/
? - kubectl get svc -n ingress-nginx
? - kubectl logs -n ingress-nginx -l app.kubernetes.io/name=ingress-nginx
? - getent hosts greenlight.local
? - minikube tunnel

enable the NGINX Ingress controller on Minikube
`minikube addons enable ingress`
- get the ingress controller status
  `kubectl get pods -n ingress-nginx`

`kubectl get ingress`

`kubectl get svc`

`kubectl port-forward svc/greenlight-api 4000:4000`

- get the IP
 `minikube ip`
Check ingress is set up properly
`kubectl describe ingress greenlight-api-ingress`
Check the log
`kubectl logs -n ingress-nginx -l app.kubernetes.io/name=ingress-nginx`

Run a test pod:
`kubectl run curlpod --image=busybox:1.28 --restart=Never -it --rm -- sh`
`wget -qO- http://greenlight-api.default.svc.cluster.local:4000/v1/healthcheck`

`kubectl exec -it -n ingress-nginx deployment/ingress-nginx-controller -- cat /etc/nginx/nginx.conf`

enable debug logs for the NGINX Ingress Controller:
`kubectl create configmap -n ingress-nginx ingress-nginx-controller --from-literal=error-log-level=debug --dry-run=client -o yaml | kubectl apply -f -`

reload the Ingress Controller:


How to get into ingress controller 
`kubectl exec -it -n ingress-nginx deploy/ingress-nginx-controller -- bash`

`kubectl exec -it -n ingress-nginx deploy/ingress-nginx-controller -- cat /etc/nginx/nginx.conf | grep -A 20 "server_name greenlight.kube"`

So far 
-----

 - I am on the `7 - Deploy REST API & its dependent services in K8s`
   - https://one2n.io/sre-bootcamp/sre-bootcamp-exercises/7-deploy-rest-api-its-dependent-services-in-k8s
 - Create namespace.yaml
 - Working on vault.yaml
https://claude.ai/chat/7f9bbf94-ff1b-439a-acff-7c2072ac46a3

What is `ServiceAccount`, `ClusterRole`, `ClusterRoleBinding `?

https://www.youtube.com/watch?v=MuX3m149FpI
