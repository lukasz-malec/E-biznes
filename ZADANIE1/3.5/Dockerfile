FROM ubuntu:24.04

ENV KOTLIN_VERSION=1.9.25
ENV JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
ENV PATH=$PATH:$JAVA_HOME/bin:/opt/kotlin/bin


RUN apt-get update && apt-get install -y \
    openjdk-8-jdk curl unzip && rm -rf /var/lib/apt/lists/*

RUN curl -L "https://github.com/JetBrains/kotlin/releases/download/v${KOTLIN_VERSION}/kotlin-compiler-${KOTLIN_VERSION}.zip" \
    -o /tmp/kotlin-compiler.zip \
    && unzip /tmp/kotlin-compiler.zip -d /opt \
    && mv /opt/kotlinc /opt/kotlin \
    && rm /tmp/kotlin-compiler.zip


WORKDIR /app


CMD ["/bin/bash"]