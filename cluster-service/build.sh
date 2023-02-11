SPARK_VERSION="2.4.5"
HADOOP_VERSION="2.7"

# -- Building the Images

docker build \
  -f cluster-base.Dockerfile \
  -t cluster-base .

docker build \
  --build-arg spark_version="${SPARK_VERSION}" \
  --build-arg hadoop_version="${HADOOP_VERSION}" \
  -f spark-base.Dockerfile \
  -t spark-base .

docker build \
  -f spark-master.Dockerfile \
  -t spark-master .

docker build \
  -f spark-worker.Dockerfile \
  -t spark-worker .

docker build \
  --build-arg spark_version="${SPARK_VERSION}" \
  -f spark-server.Dockerfile  --progress=plain \
  -t spark-server .

# Local copy of Notebooks and job-submit scripts outside Git change tracking
mkdir -p ./local/server
cp -R ./server/* ./local/server
