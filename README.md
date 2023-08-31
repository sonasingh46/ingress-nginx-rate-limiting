## Deploy Kind with Ingress

https://kind.sigs.k8s.io/docs/user/ingress/

## Deploy services

Build 2 docker images 
- One using the code as it is present in this repo. ( Call it service1 )
- Second just change the reponse of the endpoint to something else e.g. `Pong-New` at line no 44 in main.go ( Call it service 2)

Deploy the services using the YAML present in the config folder. Make sure to do 2 deployment and put the docker image correctly as it can be different for you. 

curl localhost/foo/ping --> Should Return Pong from service1
curl localhost/bar/ping --> Should Return Pong-new from service2

## Load Test

Load test using fortio https://github.com/fortio/fortio

Example: 

fortio load -c 2 --qps 20  -t 5s  localhost/foo/ping

-c 2 : Meaning, 2 concurrent connections
-qps 20 : Try to send 10 query per second. 
-t 5s: Do it for 5 seconds.

