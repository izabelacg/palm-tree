# palm-tree

### Deploy Zipkin

* Use [monitoring-tracing-zipkin-in-mem](monitoring-tracing-zipkin-in-mem.yaml) yaml.
* Comment / Uncomment lines depending on if _istio-system_ namespace is present.

### Redirect traces in Istio to Wavefront

* Use [zipkin-svc-redirect](zipkin-svc-redirect.yml) yaml.
* See [documentation](https://github.com/wavefrontHQ/wavefront-kubernetes/tree/master/istio).

### Deploy a Wavefront Proxy in Kubernetes

* Use [wavefront](wavefront.yml) yaml.
* See [documentation](https://docs.wavefront.com/kubernetes.html#step-1-deploy-a-wavefront-proxy-in-kubernetes).
* Instances:
    * demo.wavefront.com [Token]()
    * vmware.wavefront.com [Token](https://vmware.wavefront.com/proxies/add)

### HelloWorld App

#### Docker - Build

```sh
cd helloword-go
docker build -t izabelacg/kn-hello-direct -f Dockerfile .
```

#### Docker - Push changes

```sh
docker push izabelacg/kn-hello-direct
```