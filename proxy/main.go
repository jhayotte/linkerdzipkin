package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
)

var tracer opentracing.Tracer

const (
	serviceName        = "proxy"
	hostPort           = "127.0.0.1:8080"
	zipkinHTTPEndpoint = "http://zipkin:9411/api/v1/spans"
	debug              = false
	sameSpan           = true
	traceID128Bit      = true
)

func middlewareTracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Start tracing")

		// Loop through headers
		for name, headers := range r.Header {
			name = strings.ToLower(name)
			for _, h := range headers {
				log.Printf("%v: %v", name, h)
			}
		}

		// Try to join to a trace propagated in `req`.
		wireContext, err := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			fmt.Printf("error encountered while trying to extract span: %+v\n", err)
		}

		// create span
		span := tracer.StartSpan("proxy forward", ext.RPCServerOption(wireContext))
		span.SetTag("MyKEY", "specificvalue")
		defer span.Finish()

		// store span in context
		ctx := opentracing.ContextWithSpan(r.Context(), span)

		// update request context to include our new span
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
		log.Println("End Tracing")
	})
}

func forward(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(fmt.Sprintf("http://linkerd_proxy:8070/%s", r.URL.Path[1:]))
	if err != nil {
		log.Println(w, "Proxy: %s \n error: %s", r.URL.Path[1:], err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, "Proxy: root %s \n%s", r.URL.Path[1:], string(body))
}

func main() {
	collector, err := zipkin.NewHTTPCollector(zipkinHTTPEndpoint)
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v", err)
		os.Exit(-1)
	}

	// create recorder.
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	// create tracer.
	tracer, err = zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(sameSpan),
		zipkin.TraceID128Bit(traceID128Bit),
	)

	log.Println("server listening on", hostPort)
	forwardHandler := http.HandlerFunc(forward)
	http.Handle("/", middlewareTracing(forwardHandler))
	http.ListenAndServe(":8080", nil)
}
