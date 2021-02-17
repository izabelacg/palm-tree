### Metrics pipelines

#### wf-metrics-send-request

Send requests to Serverless applications periodically. It was created for testing purposes and it is being used to observe metrics going through Wavefront and to build dashboards.

- K8s Cronjob

```
kubectl logs -n wf-metrics $(kubectl get pods -n wf-metrics -l metric-send-request --sort-by=.metadata.creationTimestamp -o 'jsonpath={.items[-1].metadata.name}')

watch "kubectl get pods -n wf-metrics -l metric-send-request --sort-by=.metadata.creationTimestamp -o 'jsonpath={.items[-1].metadata.name}' | xargs kubectl logs -n wf-metrics"
```

#### wf-metrics-finder

It periodically runs a script to search Wavefront for specific metrics from Serverless. It errors out if a metric is not found in the last 20 minutes or so. It was created for testing purposes to mainly monitor the activator pod that from time to time stops emitting metrics for an unknown reason.

- K8s Cronjob

```
kubectl logs -n wf-metrics $(kubectl get pods -n wf-metrics -l metric-finder --sort-by=.metadata.creationTimestamp -o 'jsonpath={.items[-1].metadata.name}')

watch "kubectl get pods -n wf-metrics -l metric-finder --sort-by=.metadata.creationTimestamp -o 'jsonpath={.items[-1].metadata.name}' | xargs kubectl logs -n wf-metrics"
```