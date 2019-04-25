# in-toto-kubectl

This is a kubectl plugin to run in-toto verification on the images in your
kubernetes pods.

## Install

run `make deploy` and the plugin should be installed to `~/.kube/plugins`. You
can change the target by changing the KUBEPATH environment variable. For
example `make deploy KUBEPATH=~/bin` will install it to a user-controlled
`bin/` folder.

## Usage

Make sure the plugin executable was installed to somewhere in your `$PATH`, or
to add `~/.kube/plugins` to your path. Afterwards, you can use it within
kubectl:

```bash
kubectl in-toto pod/[podname]
```

In order to scan a pod, you'd have to have the link metadata and the layout in
your current folder. After passing the pod/podname argument, you can also use
`-k` and `-l` in the same way as `in-toto-verify` to pass key and layout
parameters.

### Extensions

The kubectl plugin uses parameter substitution to provide you with a
`{IMAGE_ID}` parameter that you can substitute inside of your layouts. 

In addition, a file (if it doesn't exist) called `image_id` will be populated
on the directory when verification starts. This can be used to e.g., verify
against the output of `docker build`. This second extension will disappear in
future releases, and once resource type identifiers are provided by the in-toto
framework.


## Example

An example repository exists under the `example` directory. It contains all the
tools you need to create a layout (using the python implementation), create
signed metadata files (you will need docker to build the container). If you're
using minikube to run the example, I also suggest you expose the Docker socket
before executing the functionary step so as to create the image inside the
container.

# Credit

This was very heavily based off of [stefanprodan's kubectl-kubesec plugin](https://github.com/stefanprodan/kubectl-in-toto)
