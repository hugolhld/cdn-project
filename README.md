# cdn-project - Hugo Lhernould - Groupe 13

## Requierements

- Docker
- Minikube
- Kuberneties

## To start with Minikube

- Launch minikube with `minikube start --driver=docker`
- Install Ingress with `minikube addons enable ingress`
- Build your front and back end image with docker:
    `
    cd backend
    docker build -t cdn-backend
    `
    `
    cd ../frontend
    docker build -t cdn-frontend
    `
- Load your image in minikube 
    - `minikube image load cdn-backend`
    - `minikube image load cdn-frontend`
- Applys config `kubeclt apply -f k8s/`
- Launch a `minikube tunnel` and let it run in the background
- Run `kubectl get services` and copy external ip of backend
- Run `echo YOUR_EXTERNAL_IP backend.local | sudo tee -a /ect/hosts`
- Run `kubectl get ingress`, you should have ADDRESS field
- Finally run `minikube service frontend-service --url`
- You have your frontend url ! :tada:

## To start with docker

- At root of folder, run `docker compose up --build`
- When it's done, you can access with `http://localhost/` for front and `http://localhost/api` for the back
- You can access to mongodb viewer at `http://localhost:8081/`
