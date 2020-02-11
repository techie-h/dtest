# Diginex Test

## Tech Stack used:
  ```
  Minikube - version 1.6.2 (with the tiller add on enabled)
  Kubectl  - version 1.17.0
  ```

## Instructions to run the code
* Ensure minikube is running (or start it using: `minikube start`)
* In the root folder of the project (where this README file exists), run the following command `make deploy_application`. This will deploy the application on minikube
* Run this command `minikube service api-service --url` to retrieve the base URL of the application

## Instructions to tear down the application
* In the root folder of the project (where this README file exists), run the following command `make delete_all`.

## Test application via curl
* Replace the <url> in the below line with the url retrieved from running `minikube service api-service --url`
* `curl -X POST -d '{"message": "now"}' <url>/api --header "Content-Type:application/json"`

## Tests
* There is 1 simple test within the api folder. This can be run with `go test -v`


## Time constraints
* Severe time constraints have meant, that I've ony done the very basic to complete the overall, so not many tests and no helm package manager :(


## Todo Go code
* In the go api code, there are some validation checks, that should really be in a helpers or utility file/package for reusability
* The port number should be derived from an environment variable, instead of being hard coded ...
* Add a unique ID to each request for tracebility for each user request across all micro services
* Telemetry / Instrumentation - to ensure performance of this code is measured


## To do K8s/Deployment side
* Implement k8s deployment via Helm3 package manager
* Add security policies in the YAML file for Helm etc


## Part 3 Answers
* I would deploy this application using EKS (managed Kubernetes on AWS) and Terraform module in TF version 0.12 - to set up the EKS cluster
* Using EKS - I would use Helm3 alongside adding Docker images to ECR ( Elastic Container Registry). 
* You could promote the image from Dev through to Production, tagging as required within the CI/CD process
* Each application in k8s should be in it's own namespace, with granular permissions as required
* I would use Helm 3 to deploy a chart(s) of application(s)
* I love Helm especially the fact compared to version 2, as setting/ creating certs in Helm 2 is not that great etc
* To avoid down time, I would use a Blue/Green method of deploying. This essentially means, having 2 identical production environments called blue and green
* TO avoid down time, you should deploy to the prod environment that is not live and once happy the tests have passed, switch the routing to the non live colour
* CI/CD should be used to run integration, unit and smoke, security checks (Docker image, application security etc) and linting tests while continously deploying

