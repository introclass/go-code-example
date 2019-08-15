IMAGE=envoyproxy/envoy:v1.11.0
docker run -idt --network=host -v `pwd`/envoy-0-example.yaml:/etc/envoy/envoy.yaml -v `pwd`/log:/var/log/envoy $IMAGE
