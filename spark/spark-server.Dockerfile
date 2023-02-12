FROM cluster-base

# -- Layer: JupyterLab

ARG spark_version=2.4.5
ARG jupyterlab_version=2.1.5


RUN apt-get update -y && \
    apt-get install -y python3-pip && \
    pip3 install wget pyspark==${spark_version} jupyterlab==${jupyterlab_version}

RUN apt-get install libbz2-dev -y

RUN pip3 install --upgrade pip && \
    pip3 install pandas kafka-python
# -- Runtime

EXPOSE 8888
WORKDIR ${SHARED_WORKSPACE}
CMD jupyter lab --ip=0.0.0.0 --port=8888 --no-browser --allow-root --NotebookApp.token=
