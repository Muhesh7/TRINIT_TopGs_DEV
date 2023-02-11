FROM spark-base

# -- Layer: JupyterLab

ARG spark_version=2.4.5

RUN apt-get update -y && \
    apt-get install -y python3-pip && \
    pip install --upgrade pip \
    pip install pyspark==${spark_version} grpcio grpcio-tools flask && \
    rm -rf /var/lib/apt/lists/* && \
    ln -s /usr/local/bin/python3 /usr/bin/python

# -- Runtime

EXPOSE 8888
WORKDIR ${SHARED_WORKSPACE}

ENTRYPOINT ./server/spark-submit.sh || tail -f /dev/null

