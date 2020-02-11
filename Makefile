deploy_api:
	@echo "Building API docker image ..."
	eval $(minikube docker-env)
	docker build ./api/ -t hrk/api-service:v1.0
	@echo "Deploying API service in kubernetes ..."
	kubectl apply -f ./deployment/api-service.yaml

deploy_reverse:
	@echo "Building Reverse docker image ..."
	eval $(minikube docker-env)
	docker build ./reverse/ -t hrk/reverse-service:v1.0
	@echo "Deploying Reverse service in kubernetes ..."
	kubectl apply -f ./deployment/reverse-service.yaml

deploy_application: deploy_api deploy_reverse
	@echo "Deployed Diginex application ..."

delete_api:
	@echo "Deleting api-service ..."
	kubectl delete svc api-service
	@echo "Deleting api-deployment ..."
	kubectl delete deployment api-deployment

delete_reverse:
	@echo "Deleting reverse-service ..."
	kubectl delete svc reverse-service
	@echo "Deleting reverse-deployment ..."
	kubectl delete deployment reverse-deployment

delete_all: delete_api delete_reverse
	@echo "Deleted all the resources created"


