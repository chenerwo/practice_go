package main

import (
	"mytest/grpc/pack"
)

func main()  {
	/*
	rpc := grpc.NewServer()
	Service_Proto.RegisterProdServiceServer(rpc, new(Service_Proto.ProdService))

	//lis, _ := net.Listen("tcp", ":8098");
	//rpc.Serve(lis)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		rpc.ServeHTTP(responseWriter, request)
	})
	httpServer := &http.Server{
		Addr:":8081",
		Handler:mux,
	}
	httpServer.ListenAndServe()
	 */
	pack.CurlTest()
}

