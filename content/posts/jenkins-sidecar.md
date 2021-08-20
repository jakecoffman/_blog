---
title: Jenkins sidecar
date: 2021-04-12T08:29:20-05:00
tags: [testing]
---

Jenkins has "sidecar" functionality that allows you to run multiple Docker containers at the same time to make testing easier. It seems like hardly anyone is taking advantage of this functionality, so I went ahead and pieced together how to run a Jenkins sidecar with Postgres to run a Go integration test.


<!--more-->


```groovy
stage('Build & Test') {
    steps {
        script {
            // withDockerNetwork is defined at the bottom
            withDockerNetwork { n ->
                // withRun starts postgres as a sidecar
                docker.image("postgres:13").withRun("--network ${n} -e POSTGRES_HOST_AUTH_METHOD=trust -e POSTGRES_USER=root -p 5432:5432") { c ->
                    // extract the IP address from the postgres container
                    String postgres_ip = sh(
                            script: "docker inspect -f {{.NetworkSettings.Networks.${n}.IPAddress}} ${c.id}",
                            returnStdout: true)

                    // inside means the sh will be run inside the container
                    docker.image("golang:1.16").inside("--network ${n}") {
                        // now inject the IP into an ENV and run the tests
                        sh """
export POSTGRES_IP="${postgres_ip.replace("\n", "")}"
go test -race ./...
"""
                    }
                }
            }
        }
    }
}

// withDockerNetwork creates a docker network so the containers can talk to eachother
def withDockerNetwork(Closure inner) {
    try {
        networkId = generator(9)
        sh "docker network create ${networkId}"
        inner.call(networkId)
    } finally {
        sh "docker network rm ${networkId}"
    }
}

// since networks need to be unique we generate a name that way two runs can be in parallel
def generator(int n) {
    char[] chars = "abcdefghijklmnopqrstuvwxyz".toCharArray();
    StringBuilder sb = new StringBuilder(n);
    Random random = new Random();
    for (int i = 0; i < n; i++) {
        char c = chars[random.nextInt(chars.length)];
        sb.append(c);
    }
    return sb.toString()
}

```
