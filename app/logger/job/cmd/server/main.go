package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/tx7do/kratos-transport/transport/kafka"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSdk "go.opentelemetry.io/otel/sdk/trace"
	semConv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-cqrs/app/logger/job/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     = "kratos.logger.job"
	Version  = "1.0.0"
	flagConf string
	Id, _    = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, ks *kafka.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(Id+"."+Name),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			ks,
		),
		kratos.Registrar(rr),
	)
}

func NewTracerProvider(conf *conf.Trace) error {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Endpoint)))
	if err != nil {
		return err
	}

	tp := traceSdk.NewTracerProvider(
		traceSdk.WithSampler(traceSdk.ParentBased(traceSdk.TraceIDRatioBased(1.0))),
		traceSdk.WithBatcher(exp),
		traceSdk.WithResource(resource.NewSchemaless(
			semConv.ServiceNameKey.String(Name),
			semConv.ServiceVersionKey.String(Version),
			semConv.ServiceInstanceIDKey.String(Id),
			attribute.String("env", "dev"),
		)),
	)

	otel.SetTracerProvider(tp)

	return nil
}

func NewLoggerProvider() log.Logger {
	return log.With(
		log.NewStdLogger(os.Stdout),
		"service.Id", Id,
		"service.name", Name,
		"service.version", Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

func loadConfig() (*conf.Bootstrap, *conf.Registry) {
	c := config.New(
		config.WithSource(
			file.NewSource(flagConf),
		),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}

	return &bc, &rc
}

func main() {
	flag.Parse()

	logger := NewLoggerProvider()

	bc, rc := loadConfig()
	if bc == nil || rc == nil {
		panic("load config failed")
	}

	err := NewTracerProvider(bc.Trace)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(bc.Server, rc, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
