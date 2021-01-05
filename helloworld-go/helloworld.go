package main

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/wavefronthq/wavefront-opentracing-sdk-go/reporter"
	"github.com/wavefronthq/wavefront-opentracing-sdk-go/tracer"
	"github.com/wavefronthq/wavefront-sdk-go/application"
	"github.com/wavefronthq/wavefront-sdk-go/senders"

	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("helloworld: received a request")

	// From the example https://github.com/opentracing/opentracing-go

	var serverSpan opentracing.Span
	appSpecificOperationName := "print-hello"
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		// Optionally record something about err here
		log.Printf("Couldn't extract information from headers: %s\n", err.Error())
	}

	// Create the span referring to the RPC client if available.
	// If wireContext == nil, a root span will be created.
	serverSpan = opentracing.StartSpan(
		appSpecificOperationName,
		ext.RPCServerOption(wireContext))

	defer serverSpan.Finish()

	opentracing.ContextWithSpan(context.Background(), serverSpan)

	// -- end

	time.Sleep(time.Second * 30)

	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", target)
}

func main() {
	log.Print("helloworld: starting server...")

	NewGlobalTracer("kn-hello-app", "kn-hello-service")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func NewGlobalTracer(applicationName string, serviceName string) io.Closer {

	// From: https://demo.wavefront.com/integration/go/setup
	// and https://github.com/wavefrontHQ/wavefront-opentracing-sdk-go

	directCfg := directConfiguration()
	sender, err := senders.NewDirectSender(directCfg)
	if err != nil {
		panic(err)
	}

	appTags     := application.New(applicationName, serviceName)
	wfReporter := reporter.New(sender, appTags, reporter.Source("hello-tracing.serverless-demo.example.com"))
	clReporter := reporter.NewConsoleSpanReporter("hello-tracing.serverless-demo.example.com") //Specify the same source you used for the WavefrontSpanReporter
	reporter := reporter.NewCompositeSpanReporter(wfReporter, clReporter)
	tracer      := tracer.New(reporter)

	opentracing.SetGlobalTracer(tracer)

	return ioutil.NopCloser(nil)

}

func directConfiguration() *senders.DirectConfiguration {
	return &senders.DirectConfiguration{
		Server: "https://vmware.wavefront.com",
		Token:  "3d67b8df-7e10-4365-8be7-3c526c888686",
	}
}

}