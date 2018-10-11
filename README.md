# in-toto-kubesec

This is a kubectl plugin to run in-toto verification on the images in your
kubernetes pods.

### Install

Run `make` and the plugin will be built and installed  into your .kube/plugins/
folder

### Usage

Scan a Pod:

```bash
kubectl plugin in-toto pod/[podname]
```

### Credit

This was very heavily based off of [stefanprodan's kubectl-kubesec plugin](https://github.com/stefanprodan/kubectl-in-toto)
