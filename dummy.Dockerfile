FROM ubuntu:19.10

RUN echo "echo & sleep 10000" > docker_entrypoint.sh

RUN chmod +x docker_entrypoint.sh

ENTRYPOINT ["/usr/bin/bash", "--", "docker_entrypoint.sh"]
