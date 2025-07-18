// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package weaver

import (
	"os"

	"github.com/thunur/weaver/internal/traceio"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

// tracer returns a tracer for the provided app, deploymentId, and weaveletId
// that uses the provided exporter. The tracer is also set as the otel default.
//
// Note that we set the ServiceNameKey attribute to the name of the app. Based on
// the otel resource definition in [1], the ServiceNameKey should be unique across
// all instances of the application. Ideally, it should be the name of the
// colocation group, however, given that the trace provider is globally defined
// per application, it makes sense to be the application name.
//
// [1] https://github.com/open-telemetry/opentelemetry-go/blob/v1.20.0/semconv/v1.7.0/resource.go#L813
func tracer(exporter sdktrace.SpanExporter, app, deploymentId, weaveletId string) trace.Tracer {
	const instrumentationLibrary = "github.com/thunur/weaver/serviceweaver"
	const instrumentationVersion = "0.0.1"
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(app),
			semconv.ServiceInstanceIDKey.String(weaveletId),
			semconv.ProcessPIDKey.Int(os.Getpid()),
			traceio.AppTraceKey.String(app),
			traceio.DeploymentIdTraceKey.String(deploymentId),
			traceio.WeaveletIdTraceKey.String(weaveletId),
		)),
	)
	tracer := tracerProvider.Tracer(instrumentationLibrary, trace.WithInstrumentationVersion(instrumentationVersion))

	// Set global tracing defaults.
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tracer
}
