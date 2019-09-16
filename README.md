# Envoy Proxy playground

Learning using this as the sandbox: https://github.com/envoyproxy/envoy/tree/master/examples/front-proxy

Run: `docker-compose up -d --scale service1=3 --scale service2=3`

Url endpoints:

* `http://localhost:8001/`
* `http://localhost:8081/`
* `http://localhost:8000/service/1`
* `http://localhost:8000/service/2`

## Build container

`docker build -f Dockerfile-service -t registry.gitlab.com/joekr/envoy_proxy_playground:service1 .`
`docker build -f Dockerfile-service2a -t registry.gitlab.com/joekr/envoy_proxy_playground:service2a .`
`docker build -f Dockerfile-service2b -t registry.gitlab.com/joekr/envoy_proxy_playground:service2b .`

`docker login registry.gitlab.com`
`docker push registry.gitlab.com/joekr/envoy_proxy_playground:service1`
`docker push registry.gitlab.com/joekr/envoy_proxy_playground:service2a`
`docker push registry.gitlab.com/joekr/envoy_proxy_playground:service2b`

## run locally

Update envoy proxy `dc restart edge`

## Goal

Launch Infrastructure which supports using Envoy Proxy to route canary requests to new cluster. If new cluster is passing canary tests, then move into production.


## Other sources
- https://blog.envoyproxy.io/envoy-hot-restart-1d16b14555b5

### Cloud options

- https://docs.microsoft.com/en-us/azure/devops/migrate/phase-rollout-with-rings?view=azure-devops
- https://aws.amazon.com/blogs/compute/implementing-canary-deployments-of-aws-lambda-functions-with-alias-traffic-shifting/
