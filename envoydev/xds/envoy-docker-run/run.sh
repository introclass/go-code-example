IMAGE=lijiaocn/envoy:v1.11.0
NAME=envoy-1.11.0

if  [ $# -ne 1 ];then
	echo "must choose one config file"
	exit 1
fi

config=$1

docker rm -f $NAME 2>&1 1>/dev/null

if [ `uname`="Darwin" ];then
#	docker run -idt --name $NAME -e loglevel=debug -p 9901:9901 -p 80-86:80-86 -v `pwd`/$config:/etc/envoy/envoy.yaml -v `pwd`/log:/var/log/envoy $IMAGE
	docker run -idt --name $NAME -p 9901:9901 -p 80-86:80-86 -v `pwd`/$config:/etc/envoy/envoy.yaml -v `pwd`/log:/var/log/envoy $IMAGE
else
	docker run -idt --name $NAME -e loglevel=debug --network=host -v `pwd`/$config:/etc/envoy/envoy.yaml -v `pwd`/log:/var/log/envoy $IMAGE
fi
