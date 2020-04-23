# GoogleBookAPI Example

Google book API application is build on top of [flogo](https://www.flogo.io/) application to demostrate zbio messaging in flow based application.
ZBIO Topic `googleBookAPI` would be created in zbio and messages are sent to those topic whenever request is made to the application to retrieve google book details.

## Commands to build and deploy Googlebookapi app in k8s cluster

### How to use in localhost

* Open http://localhost:9999/books/9788126568772
  * isbn_number: `9788126568772` ; URL: http://localhost:9999/books/<isbn_number>
* ISBN stands for international standard book number , which is
    13 digit number uniquely identify all the books.
* ISBN can be found in internet. **ISBN:** `9781788999786`, **Book Name:** `Mastering Kubernetes`
* The digit should not have any special characters in between.

### When using Kind cluster

```bash

# Build the googlebookapi application and dockerise;
# Generate image name: zbio/googlebookapi:v1
make

# If using kind cluster
kind load docker-image zbio/googlebookapi:v1

# Deploy application in Kubernetes
make deploy
```

### Cleaning

```bash
# Deletes googlebookapi binaries
# Delete deployed application from Kubernetes
make clean
```