# palm-tree

### Deploy Zipkin

* Use [monitoring-tracing-zipkin-in-mem](monitoring-tracing-zipkin-in-mem.yaml) yaml.
* Comment / Uncomment lines depending on if _istio-system_ namespace is present.

### Redirect traces in Istio to Wavefront

* Use [zipkin-svc-redirect](zipkin-svc-redirect.yml) yaml.
* See [documentation](https://github.com/wavefrontHQ/wavefront-kubernetes/tree/master/istio).