// Create: 2018/12/29 18:32:00 Change: 2018/12/29 18:32:00
// FileName: main.go
// Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
//
// Distributed under terms of the GPL license.

package main

import (
	"fmt"
	"net"
	"time"

	api_v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/listener"
	"github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	http_router "github.com/envoyproxy/go-control-plane/envoy/config/filter/http/router/v2"
	http_conn_manager "github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	xds "github.com/envoyproxy/go-control-plane/pkg/server"
	"github.com/envoyproxy/go-control-plane/pkg/util"
	proto_type "github.com/gogo/protobuf/types"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type NodeConfig struct {
	node      *core.Node
	endpoints []cache.Resource //[]*api_v2.ClusterLoadAssignment
	clusters  []cache.Resource //[]*api_v2.Cluster
	routes    []cache.Resource //[]*api_v2.RouteConfiguration
	listeners []cache.Resource //[]*api_v2.Listener
}

//implement cache.NodeHash
func (n NodeConfig) ID(node *core.Node) string {
	return node.GetId()
}

func ADD_Cluster_With_Static_Endpoint(n *NodeConfig) {

	// 每个 route 对应一个或一组 cluster，可以设置 cluster 的权重
	// 每个 cluster 包含多个 endpoints 分组(locality)，可以设置分组的权重
	// 每个 endpoints 分组内包含多个 endpoint，可以设置 endpoint 的权重

	timeout := 1 * time.Second

	// endpoint 地址
	hostIdentifier := &endpoint.LbEndpoint_Endpoint{
		Endpoint: &endpoint.Endpoint{
			Address: &core.Address{
				Address: &core.Address_SocketAddress{
					SocketAddress: &core.SocketAddress{
						Protocol: core.TCP,
						Address:  "172.16.129.26",
						PortSpecifier: &core.SocketAddress_PortValue{
							PortValue: 80,
						},
					},
				},
			},
		},
	}

	// 一个 endpoint
	lbEndpoint := &endpoint.LbEndpoint{
		HostIdentifier: hostIdentifier,
	}

	lbEndpoints := make([]*endpoint.LbEndpoint, 0)
	lbEndpoints = append(lbEndpoints, lbEndpoint)

	// endpoints 分组，由多个 endpoint 组成
	localityLbEndpoints := &endpoint.LocalityLbEndpoints{
		LbEndpoints: lbEndpoints,
	}

	endpoints := make([]*endpoint.LocalityLbEndpoints, 0)
	endpoints = append(endpoints, localityLbEndpoints)

	//cluster 的多个 endpoints 分组
	clusterLoadAssignment := &api_v2.ClusterLoadAssignment{
		ClusterName: "none",
		Endpoints:   endpoints,
	}

	// 使用静态 endpoints 的 cluster，类型为 v2.Cluster_STATIC
	cluster := &api_v2.Cluster{
		Name:        "cluster_with_static_endpoint",
		AltStatName: "cluster_with_static_endpoint",
		ClusterDiscoveryType: &api_v2.Cluster_Type{
			Type: api_v2.Cluster_STATIC,
		},
		EdsClusterConfig:              nil,
		ConnectTimeout:                &timeout,
		PerConnectionBufferLimitBytes: nil, //default 1MB
		LbPolicy:                      api_v2.Cluster_ROUND_ROBIN,
		LoadAssignment:                clusterLoadAssignment,
	}

	n.clusters = append(n.clusters, cluster)
}

func ADD_Cluster_With_Dynamic_Endpoint(n *NodeConfig) {

	timeout := 1 * time.Second

	// endpoint 地址
	hostIdentifier := &endpoint.LbEndpoint_Endpoint{
		Endpoint: &endpoint.Endpoint{
			Address: &core.Address{
				Address: &core.Address_SocketAddress{
					SocketAddress: &core.SocketAddress{
						Protocol: core.TCP,
						Address:  "182.16.129.26",
						PortSpecifier: &core.SocketAddress_PortValue{
							PortValue: 80,
						},
					},
				},
			},
		},
	}

	// 一个 endpoints
	lbEndpoint := &endpoint.LbEndpoint{
		HostIdentifier: hostIdentifier,
	}

	lbEndpoints := make([]*endpoint.LbEndpoint, 0)
	lbEndpoints = append(lbEndpoints, lbEndpoint)

	// 一个 endpoint 分组
	localityLbEndpoint := &endpoint.LocalityLbEndpoints{
		LbEndpoints: lbEndpoints,
	}

	localityLbEndpoints := make([]*endpoint.LocalityLbEndpoints, 0)
	localityLbEndpoints = append(localityLbEndpoints, localityLbEndpoint)

	// 请求量过载时的处理策略
	dropOverLoad := &api_v2.ClusterLoadAssignment_Policy_DropOverload{
		Category: "drop_policy1",
		DropPercentage: &envoy_type.FractionalPercent{
			Numerator:   3,
			Denominator: envoy_type.FractionalPercent_HUNDRED,
		},
	}

	dropOverLoads := make([]*api_v2.ClusterLoadAssignment_Policy_DropOverload, 0)
	dropOverLoads = append(dropOverLoads, dropOverLoad)

	// cluster 的多个 endpoints 分组
	point := &api_v2.ClusterLoadAssignment{
		ClusterName: "cluster_with_dynamic_endpoint",
		Endpoints:   localityLbEndpoints,
		Policy: &api_v2.ClusterLoadAssignment_Policy{
			DropOverloads: dropOverLoads,
			OverprovisioningFactor: &proto_type.UInt32Value{
				Value: 140,
			},
		},
	}

	// 写入 eds，clsuter 从 eds 中获取对应的 endpoints
	n.endpoints = append(n.endpoints, point)

	// grpc 服务地址，这里的 xds_cluster 是在 envoy 中配置的 cluster，用于发现 endpoints
	grpcService := &core.GrpcService{
		TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
			EnvoyGrpc: &core.GrpcService_EnvoyGrpc{
				ClusterName: "xds_cluster",
			},
		},
	}

	grpcServices := make([]*core.GrpcService, 0)
	grpcServices = append(grpcServices, grpcService)

	// eds 发现配置
	edsClusterConfig := &api_v2.Cluster_EdsClusterConfig{
		EdsConfig: &core.ConfigSource{
			ConfigSourceSpecifier: &core.ConfigSource_ApiConfigSource{
				ApiConfigSource: &core.ApiConfigSource{
					ApiType:      core.ApiConfigSource_GRPC,
					GrpcServices: grpcServices,
				},
			},
		},
	}

	// 通过 eds 发现 endpoint 中的 cluster，类型为 Cluster_EDS
	cluster := &api_v2.Cluster{
		Name:        "cluster_with_dynamic_endpoint",
		AltStatName: "cluster_with_dynamic_endpoint",
		ClusterDiscoveryType: &api_v2.Cluster_Type{
			Type: api_v2.Cluster_EDS,
		},
		EdsClusterConfig: edsClusterConfig,
		ConnectTimeout:   &timeout,
		LbPolicy:         api_v2.Cluster_ROUND_ROBIN,
	}

	n.clusters = append(n.clusters, cluster)
}

func ADD_Listener_With_Static_Route(n *NodeConfig) {

	// listener 主要由 监听地址 和 多个 filter 组成
	// 其中 filter 是最复杂的部分，它由多条 filter 链组成
	// 名为 HttpConnectionManager 的 network filter 中继续包含有 http filter

	// 到达监听地址的请求，经过多个filter处理，最终转发到给对应的 cluster

	// listener 的监听地址
	address := &core.Address{
		Address: &core.Address_SocketAddress{
			SocketAddress: &core.SocketAddress{
				Protocol: core.TCP,
				Address:  "0.0.0.0",
				PortSpecifier: &core.SocketAddress_PortValue{
					PortValue: 9000,
				},
			},
		},
	}

	// 一个 http route， 符合条件的请求转给对应的 cluster
	rt := &route.Route{
		Match: &route.RouteMatch{
			PathSpecifier: &route.RouteMatch_Prefix{
				Prefix: "/",
			},
			CaseSensitive: &proto_type.BoolValue{
				Value: false,
			},
		},
		Action: &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{
					Cluster: "cluster_with_static_endpoint", //转发到这个cluster
				},
				HostRewriteSpecifier: &route.RouteAction_HostRewrite{
					HostRewrite: "webshell.com",
				},
			},
		},
	}

	routes := make([]*route.Route, 0)
	routes = append(routes, rt)

	// route 在 virutualhost 中，virtualhost 指定了域名
	virtualHost := &route.VirtualHost{
		Name: "local",
		Domains: []string{
			"webshell.com",
		},
		Routes: routes,
	}

	virtualHosts := make([]*route.VirtualHost, 0)
	virtualHosts = append(virtualHosts, virtualHost)

	// 准备一个 http filter
	http_filter_router_ := &http_router.Router{
		DynamicStats: &proto_type.BoolValue{
			Value: true,
		},
	}

	http_filter_router, err := util.MessageToStruct(http_filter_router_)
	if err != nil {
		glog.Error(err)
		return
	}

	httpFilter := &http_conn_manager.HttpFilter{
		Name: "envoy.router",
		ConfigType: &http_conn_manager.HttpFilter_Config{
			Config: http_filter_router,
		},
	}

	httpFilters := make([]*http_conn_manager.HttpFilter, 0)
	httpFilters = append(httpFilters, httpFilter)

	// 包含 route 的 virtualhost 和 httpfilter 汇聚到 listener_filter 中
	listen_filter_http_conn_ := &http_conn_manager.HttpConnectionManager{
		StatPrefix: "ingress_http",
		RouteSpecifier: &http_conn_manager.HttpConnectionManager_RouteConfig{
			RouteConfig: &api_v2.RouteConfiguration{
				Name:         "None",
				VirtualHosts: virtualHosts,
			},
		},
		HttpFilters: httpFilters,
	}

	// ptypes.MarshalAny()
	listen_filter_http_conn, err := util.MessageToStruct(listen_filter_http_conn_)
	if err != nil {
		glog.Error(err)
		return
	}

	// listen_filter 被纳入最终的 filter
	filter := &listener.Filter{
		Name: "envoy.http_connection_manager",
		ConfigType: &listener.Filter_Config{
			Config: listen_filter_http_conn,
		},
	}

	filters := make([]*listener.Filter, 0)
	filters = append(filters, filter)

	filterChain := &listener.FilterChain{
		Filters: filters,
	}

	filterChains := make([]*listener.FilterChain, 0)
	filterChains = append(filterChains, filterChain)

	// 一个 listener
	lis := &api_v2.Listener{
		Name:         "listener_with_static_route_port_9000",
		Address:      address,
		FilterChains: filterChains,
	}

	n.listeners = append(n.listeners, lis)
}

func ADD_Listener_With_Dynamic_Route(n *NodeConfig) {

	rt := &route.Route{
		Match: &route.RouteMatch{
			PathSpecifier: &route.RouteMatch_Prefix{
				Prefix: "/",
			},
			CaseSensitive: &proto_type.BoolValue{
				Value: false,
			},
		},
		Action: &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{
					Cluster: "cluster_with_dynamic_endpoint",
				},
				HostRewriteSpecifier: &route.RouteAction_HostRewrite{
					HostRewrite: "webshell.com",
				},
			},
		},
	}

	routes := make([]*route.Route, 0)
	routes = append(routes, rt)

	virtualHost := &route.VirtualHost{
		Name: "local",
		Domains: []string{
			"webshell.com",
		},
		Routes: routes,
	}

	virtualHosts := make([]*route.VirtualHost, 0)
	virtualHosts = append(virtualHosts, virtualHost)

	routeConfig := &api_v2.RouteConfiguration{
		Name:         "dynamic_route",
		VirtualHosts: virtualHosts,
	}

	n.routes = append(n.routes, routeConfig)

	http_filter_router_ := &http_router.Router{
		DynamicStats: &proto_type.BoolValue{
			Value: true,
		},
	}

	http_filter_router, err := util.MessageToStruct(http_filter_router_)
	if err != nil {
		glog.Error(err)
		return
	}

	listen_filter_http_conn_ := &http_conn_manager.HttpConnectionManager{
		StatPrefix: "ingress_http",
		RouteSpecifier: &http_conn_manager.HttpConnectionManager_Rds{
			Rds: &http_conn_manager.Rds{
				RouteConfigName: "dynamic_route", //绑定的RDS
				ConfigSource: &core.ConfigSource{
					ConfigSourceSpecifier: &core.ConfigSource_ApiConfigSource{
						ApiConfigSource: &core.ApiConfigSource{
							ApiType: core.ApiConfigSource_GRPC,
							GrpcServices: []*core.GrpcService{
								&core.GrpcService{
									TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
										EnvoyGrpc: &core.GrpcService_EnvoyGrpc{
											ClusterName: "xds_cluster",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		HttpFilters: []*http_conn_manager.HttpFilter{
			&http_conn_manager.HttpFilter{
				Name: "envoy.router",
				ConfigType: &http_conn_manager.HttpFilter_Config{
					Config: http_filter_router,
				},
			},
		},
	}
	listen_filter_http_conn, err := util.MessageToStruct(listen_filter_http_conn_)
	if err != nil {
		glog.Error(err)
		return
	}

	listener := &api_v2.Listener{
		Name: "listener_with_dynamic_route_port_9001",
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.TCP,
					Address:  "0.0.0.0",
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: 9001,
					},
				},
			},
		},
		FilterChains: []*listener.FilterChain{
			&listener.FilterChain{
				Filters: []*listener.Filter{
					&listener.Filter{
						Name: "envoy.http_connection_manager",
						ConfigType: &listener.Filter_Config{
							Config: listen_filter_http_conn,
						},
					},
				},
			},
		},
	}

	n.listeners = append(n.listeners, listener)
}

func ADD_Cluster_With_ADS_Endpoint(n *NodeConfig) {

	timeout := 1 * time.Second

	socketAddress := &core.SocketAddress{
		Protocol: core.TCP,
		Address:  "192.16.129.26",
		PortSpecifier: &core.SocketAddress_PortValue{
			PortValue: 80,
		},
	}

	address := &core.Address{
		Address: &core.Address_SocketAddress{
			SocketAddress: socketAddress,
		},
	}

	hostIdentifier := &endpoint.LbEndpoint_Endpoint{
		Endpoint: &endpoint.Endpoint{
			Address: address,
		},
	}

	lbEndpoint := &endpoint.LbEndpoint{
		HostIdentifier: hostIdentifier,
	}

	lbEndpoints := make([]*endpoint.LbEndpoint, 0)
	lbEndpoints = append(lbEndpoints, lbEndpoint)

	localityLbEndpoints := &endpoint.LocalityLbEndpoints{
		LbEndpoints: lbEndpoints,
	}

	endpoints := make([]*endpoint.LocalityLbEndpoints, 0)
	endpoints = append(endpoints, localityLbEndpoints)

	point := &api_v2.ClusterLoadAssignment{
		ClusterName: "ads_endpoint",
		Endpoints:   endpoints,
		Policy: &api_v2.ClusterLoadAssignment_Policy{
			DropOverloads: []*api_v2.ClusterLoadAssignment_Policy_DropOverload{
				&api_v2.ClusterLoadAssignment_Policy_DropOverload{
					Category: "drop_policy1",
					DropPercentage: &envoy_type.FractionalPercent{
						Numerator:   3,
						Denominator: envoy_type.FractionalPercent_HUNDRED,
					},
				},
			},
			OverprovisioningFactor: &proto_type.UInt32Value{
				Value: 140,
			},
		},
	}

	n.endpoints = append(n.endpoints, point)

	cluster := &api_v2.Cluster{
		Name:           "cluster_with_ads_endpoint",
		AltStatName:    "cluster_with_ads_endpoint",
		ConnectTimeout: &timeout,
		ClusterDiscoveryType: &api_v2.Cluster_Type{
			Type: api_v2.Cluster_EDS,
		},
		LbPolicy: api_v2.Cluster_ROUND_ROBIN,
		EdsClusterConfig: &api_v2.Cluster_EdsClusterConfig{
			EdsConfig: &core.ConfigSource{
				ConfigSourceSpecifier: &core.ConfigSource_Ads{
					Ads: &core.AggregatedConfigSource{}, //使用ADS
				},
			},
			ServiceName: "ads_endpoint", //与endpoint中的ClusterName对应。
		},
	}

	n.clusters = append(n.clusters, cluster)
}

func ADD_Listener_With_ADS_Route(n *NodeConfig) {

	r := &route.Route{
		Match: &route.RouteMatch{
			PathSpecifier: &route.RouteMatch_Prefix{
				Prefix: "/",
			},
			CaseSensitive: &proto_type.BoolValue{
				Value: false,
			},
		},
		Action: &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{
					Cluster: "cluster_with_ads_endpoint",
				},
				HostRewriteSpecifier: &route.RouteAction_HostRewrite{
					HostRewrite: "webshell.com",
				},
			},
		},
	}

	routes := make([]*route.Route, 0)
	routes = append(routes, r)

	virtualHost := &route.VirtualHost{
		Name: "local",
		Domains: []string{
			"ads.webshell.com",
		},
		Routes: routes,
	}

	virtualHosts := make([]*route.VirtualHost, 0)
	virtualHosts = append(virtualHosts, virtualHost)

	routeConfig := &api_v2.RouteConfiguration{
		Name:         "ads_route",
		VirtualHosts: virtualHosts,
	}

	n.routes = append(n.routes, routeConfig)

	http_filter_router_ := &http_router.Router{
		DynamicStats: &proto_type.BoolValue{
			Value: true,
		},
	}

	http_filter_router, err := util.MessageToStruct(http_filter_router_)
	if err != nil {
		glog.Error(err)
		return
	}

	httpFilter := &http_conn_manager.HttpFilter{
		Name: "envoy.router",
		ConfigType: &http_conn_manager.HttpFilter_Config{
			Config: http_filter_router,
		},
	}

	httpFilters := make([]*http_conn_manager.HttpFilter, 0)
	httpFilters = append(httpFilters, httpFilter)

	listen_filter_http_conn_ := &http_conn_manager.HttpConnectionManager{
		StatPrefix: "ingress_http",
		RouteSpecifier: &http_conn_manager.HttpConnectionManager_Rds{
			Rds: &http_conn_manager.Rds{
				RouteConfigName: "ads_route",
				ConfigSource: &core.ConfigSource{
					ConfigSourceSpecifier: &core.ConfigSource_Ads{
						Ads: &core.AggregatedConfigSource{}, //使用ADS
					},
				},
			},
		},
		HttpFilters: httpFilters,
	}

	listen_filter_http_conn, err := util.MessageToStruct(listen_filter_http_conn_)
	if err != nil {
		glog.Error(err)
		return
	}

	filter := &listener.Filter{
		Name: "envoy.http_connection_manager",
		ConfigType: &listener.Filter_Config{
			Config: listen_filter_http_conn,
		},
	}

	filters := make([]*listener.Filter, 0)
	filters = append(filters, filter)

	filterChain := &listener.FilterChain{
		Filters: filters,
	}

	filterChains := make([]*listener.FilterChain, 0)
	filterChains = append(filterChains, filterChain)

	socketAddr := &core.SocketAddress{
		Protocol: core.TCP,
		Address:  "0.0.0.0",
		PortSpecifier: &core.SocketAddress_PortValue{
			PortValue: 9002,
		},
	}

	addr := &core.Address{
		Address: &core.Address_SocketAddress{
			SocketAddress: socketAddr,
		},
	}

	lis := &api_v2.Listener{
		Name:         "listener_with_dynamic_route_port_9002",
		Address:      addr,
		FilterChains: filterChains,
	}

	n.listeners = append(n.listeners, lis)
}

func Update_SnapshotCache(s cache.SnapshotCache, n *NodeConfig, version string) {
	err := s.SetSnapshot(n.ID(n.node), cache.NewSnapshot(version, n.endpoints, n.clusters, n.routes, n.listeners))
	if err != nil {
		glog.Error(err)
	}
}

func main() {
	snapshotCache := cache.NewSnapshotCache(false, NodeConfig{}, nil)
	server := xds.NewServer(snapshotCache, nil)
	grpcServer := grpc.NewServer()
	lis, _ := net.Listen("tcp", ":5678")

	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	api_v2.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	api_v2.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	api_v2.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	api_v2.RegisterListenerDiscoveryServiceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			// error handling
		}
	}()

	node := &core.Node{
		Id:      "envoy-64.58",
		Cluster: "test",
	}

	node_config := &NodeConfig{
		node:      node,
		endpoints: []cache.Resource{}, //[]*api_v2.ClusterLoadAssignment
		clusters:  []cache.Resource{}, //[]*api_v2.Cluster
		routes:    []cache.Resource{}, //[]*api_v2.RouteConfiguration
		listeners: []cache.Resource{}, //[]*api_v2.Listener
	}

	input := ""

	fmt.Printf("Enter to update version 1: ADD_Cluster_With_Static_Endpoint")
	_, _ = fmt.Scanf("\n", &input)
	ADD_Cluster_With_Static_Endpoint(node_config)
	Update_SnapshotCache(snapshotCache, node_config, "1")
	fmt.Printf("ok")

	fmt.Printf("\nEnter to update version 2: ADD_Cluster_With_Dynamic_Endpoint")
	_, _ = fmt.Scanf("\n", &input)
	ADD_Cluster_With_Dynamic_Endpoint(node_config)
	Update_SnapshotCache(snapshotCache, node_config, "2")
	fmt.Printf("ok")

	fmt.Printf("\nEnter to update version 3: ADD_Cluster_With_ADS_Endpoint")
	_, _ = fmt.Scanf("\n", &input)
	ADD_Cluster_With_ADS_Endpoint(node_config)
	Update_SnapshotCache(snapshotCache, node_config, "3")
	fmt.Printf("ok")

	fmt.Printf("\nEnter to update version 4: ADD_Listener_With_Static_Route")
	_, _ = fmt.Scanf("\n", &input)
	ADD_Listener_With_Static_Route(node_config)
	Update_SnapshotCache(snapshotCache, node_config, "4")
	fmt.Printf("ok")

	fmt.Printf("\nEnter to update version 5: ADD_Listener_With_Dynamic_Route")
	_, _ = fmt.Scanf("\n", &input)
	ADD_Listener_With_Dynamic_Route(node_config)
	Update_SnapshotCache(snapshotCache, node_config, "5")
	fmt.Printf("ok")

	fmt.Printf("\nEnter to update version 6: ADD_Listener_With_ADS_Route")
	_, _ = fmt.Scanf("\n", &input)
	ADD_Listener_With_ADS_Route(node_config)
	Update_SnapshotCache(snapshotCache, node_config, "6")
	fmt.Printf("ok")

	fmt.Printf("\nEnter to exit: ")
	_, _ = fmt.Scanf("\n", &input)
}

//一个端口被多少个Listener使用，是否允许listener名称不同但端口相同 ？
//一个 Listener可以使用多个filter，是否支持多个同类型的filter?
//名为 envoy.http_connection_manager filter 中是否可以有多个http filter，是否支持多个同类型的http filter ？
//名为 envoy.http_connection_manager filter 中是否可以有多个 VirtualHosts ？
