IMAGE=envoyproxy/envoy:v1.11.0
#docker run -idt --network=host -v `pwd`/envoy-static.yaml:/etc/envoy/envoy.yaml -v `pwd`/log:/var/log/envoy $IMAGE
docker run -idt --network=host -v `pwd`/envoy-ads.yaml:/etc/envoy/envoy.yaml -v `pwd`/log:/var/log/envoy $IMAGE
