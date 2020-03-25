package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/cadvisor/client"
	v1 "github.com/google/cadvisor/info/v1"
	"github.com/sirupsen/logrus"
)

func GroupByPod(cinfo []v1.ContainerInfo) (map[string][]v1.ContainerInfo, error) {
	groups := make(map[string][]v1.ContainerInfo)
	for _, v := range cinfo {
		key := strings.Split(v.Name, v.Id)[0]
		_, ok := groups[key]
		if !ok {
			groups[key] = make([]v1.ContainerInfo, 1)
			groups[key][0] = v
			continue
		}
		groups[key] = append(groups[key], v)
	}
	return groups, nil
}

func main() {
	cadvisor, err := client.NewClient("http://172.29.128.20:8080")
	cinfoList, err := cadvisor.AllDockerContainers(
		&v1.ContainerInfoRequest{
			NumStats: 60,
			Start:    time.Time{},
			End:      time.Time{},
		})
	if err != nil {
		logrus.Errorf("read cadvisor fail, stop collect: %s", err.Error())
		return
	}

	//bytes, err := json.MarshalIndent(cinfoList, "", "  ")
	//if err != nil {
	//	glog.Fatal(err.Error())
	//}
	//fmt.Printf("%s\n", bytes)

	pods, _ := GroupByPod(cinfoList)
	for key, v := range pods {
		fmt.Printf("%s: %s\n", key, v[0].Aliases[0])
	}
}
