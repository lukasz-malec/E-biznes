FROM ubuntu:22.04

RUN apt update && \
    apt install -y python3.10 && \
    ln -sf /usr/bin/python3.10 /usr/bin/python3 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR "/app"

CMD ["bash"]