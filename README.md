# in-toto-kubesec

This is a kubectl plugin to run in-toto verification on the images in your
kubernetes pods.

### Install

run `make deploy` and the plugin should be installed to `~/.kube/plugins`. You
can change the target by changing the KUBEPATH environment variable. For
example `make deploy KUBEPATH=~/bin` will install it to a user-controlled
`bin/` folder.

### Usage

Make sure the plugin executable was installed to somewhere in your `$PATH`, or
to add `~/.kube/plugins` to your path. Afterwards, you can use it within
kubectl:

```bash
kubectl plugin in-toto pod/[podname]
```

### Credit

This was very heavily based off of [stefanprodan's kubectl-kubesec plugin](https://github.com/stefanprodan/kubectl-in-toto)
