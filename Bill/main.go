package main

import (
	pb "github.com/shubha-intelops/grpcx/bill/gen/api/v1"
	grpccontrollers "github.com/shubha-intelops/grpcx/bill/pkg/grpc/server/controllers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"

	"context"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	restcontrollers "github.com/shubha-intelops/grpcx/bill/pkg/rest/server/controllers"
	"github.com/sinhashubham95/go-actuator"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
)

var (
	host = "localhost"
	port = "8085"
)

func main() {

	// Set up the TCP listener
	addr := net.JoinHostPort(host, port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Errorf("error while starting TCP listener: %s", err)
		return
	}

	log.Printf("TCP listener started at port: %s", port)

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Create the Bill server
	billServer, err := grpccontrollers.NewBillServer()
	if err != nil {
		log.Errorf("error while creating billServer: %s", err)
		return
	}
	// Register the Bill server with the gRPC server
	pb.RegisterBillServiceServer(grpcServer, billServer)

	// Enable reflection for the gRPC server
	reflection.Register(grpcServer)

	go func() {
		// Start serving gRPC requests
		if err = grpcServer.Serve(lis); err != nil {
			log.Errorf("error serving gRPC: %s", err)
			return
		}
	}()

	router := gin.Default()
	if len(serviceName) > 0 && len(collectorURL) > 0 {
		// add opentel
		cleanup := initTracer()
		defer func(func(context.Context) error) {
			_ = cleanup(context.Background())
		}(cleanup)
		router.Use(otelgin.Middleware(serviceName))
	}

	// add actuator
	addActuator(router)
	// add prometheus
	addPrometheus(router)

	billController, err := restcontrollers.NewBillController()
	if err != nil {
		log.Errorf("error occurred: %s", err)
		return
	}

	v1 := router.Group("/v1")
	{

		v1.GET("/bills/:id", billController.FetchBill)
		v1.POST("/bills", billController.CreateBill)
		v1.PUT("/bills/:id", billController.UpdateBill)
		v1.DELETE("/bills/:id", billController.DeleteBill)
		v1.GET("/bills", billController.ListBills)
		v1.PATCH("/bills/:id", billController.PatchBill)
		v1.HEAD("/bills", billController.HeadBill)
		v1.OPTIONS("/bills", billController.OptionsBill)

	}

	Port := ":8084"
	log.Println("Server started")
	if err = router.Run(Port); err != nil {
		log.Errorf("error occurred: %s", err)
		return
	}

}

var (
	serviceName  = os.Getenv("SERVICE_NAME")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("INSECURE_MODE")
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func addPrometheus(router *gin.Engine) {
	router.GET("/metrics", prometheusHandler())
}

func addActuator(router *gin.Engine) {
	actuatorHandler := actuator.GetActuatorHandler(&actuator.Config{Endpoints: []int{
		actuator.Env,
		actuator.Info,
		actuator.Metrics,
		actuator.Ping,
		// actuator.Shutdown,
		actuator.ThreadDump,
	},
		Env:     "dev",
		Name:    "bill",
		Port:    8084,
		Version: "0.0.1",
	})
	ginActuatorHandler := func(ctx *gin.Context) {
		actuatorHandler(ctx.Writer, ctx.Request)
	}
	router.GET("/actuator/*endpoint", ginActuatorHandler)
}

func initTracer() func(context.Context) error {
	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorURL),
		),
	)

	if err != nil {
		log.Errorf("error occurred: %s", err)
		return nil
	}
	restResources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("services.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Printf("could not set restResources: %s", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(exporter),
			sdktrace.WithResource(restResources),
		),
	)
	return exporter.Shutdown
}

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}
