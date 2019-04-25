Functionary step
================

This is somewhat a silly example to produce a small image that can be kept in
the cluster running without actually doing anything (I'm sure there are ways to
do this without a 700k binary, but oh well). First, you need to compile the
entrypoint:

```
gcc --static main.c
```

Then build the image and create an in-toto link. To do this easily in minikube
I suggest exposing the docker socket:

```
eval $(minikube docker-env)
in-toto-run -k functionary -n build -p image_id -- docker build . -t empty:1.0.0 --iidfile image_id
```

as a side note, the `image_id` file will be used to serialize the hash of the
resulting image to a file to be tracked as an in-toto product.
