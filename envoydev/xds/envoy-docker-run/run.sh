IMAGE=envoyproxy/envoy:v1.11.0

if  [ $# -ne 1 ];then
	echo "must choose one config file"
	exit 1
fi

config=$1

if [ `uname`="Darwin" ];then
	docker run -idt -p 9901:9901 -p 80-86:80-86 -v `pwd`/$config:/etc/envoy/envoy.yaml -v `pwd`/log:/var/log/envoy $IMAGE
else
	docker run -idt --network=host -v `pwd`/$config:/etc/envoy/envoy.yaml -v `pwd`/log:/var/log/envoy $IMAGE
fi
