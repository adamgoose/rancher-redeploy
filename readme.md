# Redeploy for Rancher

Redeploy is a very simple golang CLI that utilizes the Kubernetes API to "redeploy" a service within the cluster. It takes the Rancher approach to doing so, in that it does two things to the given deployment:

- Increments the "deployment.kubernetes.io/revision" annotation numerically
- Sets the "cattle.io/timestamp" annotation of the "spec.template" to the current timestamp.

It is currently used in CI/CD pipelines by [Adeo Healthcare Software](https://adeohs.com).

## Usage

```
Redeploys a service in Rancher

Usage:
  redeploy {namespace} {service} [flags]

Flags:
  -d, --development         Use out-of-cluster config
  -c, --kubeconfig string   kubeconfig file to use while using out-of-cluster config (default "~/.kube/config")
  -h, --help                help for redeploy
```