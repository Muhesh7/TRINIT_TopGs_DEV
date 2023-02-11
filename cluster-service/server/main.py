import os

from pyspark.sql import SparkSession
#region
# app = Flask(__name__)


# def produce_pi(scale):
#     spark = SparkSession.builder.appName("PythonPi").getOrCreate()
#     n = 100000 * scale

#     def f(_):
#         from random import random
#         x = random()
#         y = random()
#         return 1 if x ** 2 + y ** 2 <= 1 else 0

#     count = spark.sparkContext.parallelize(
#         range(1, n + 1), scale).map(f).reduce(lambda x, y: x + y)
#     spark.stop()
#     pi = 4.0 * count / n
#     return pi


# @app.route("/")
# def index():
#     return "Python Flask SparkPi server running. Add the 'sparkpi' route to this URL to invoke the app."


# @app.route("/sparkpi")
# def sparkpi():
#     scale = int(request.args.get('scale', 2))
#     pi = produce_pi(scale)
#     response = "Pi is roughly {}".format(pi)

#     return response


# if __name__ == "__main__":
#     port = int(os.environ.get("PORT", 8888))
#     app.run(host='0.0.0.0', port=port)
#endregion

# import pandas as pd

# user_details_df = pd.read_csv('user_details.csv', low_memory=False)

# print(df.info())

features = ['name', 'age', 'country', 'ip']

rules = {
    # 'name': 'exact',
    'age': 'exact',
    # 'country': 'exact',
    # 'ip': 'exact',
}

clusters = []

def cluster_stream_data(row_id, row, features, rules):
    cluster_id = 0
    if len(clusters) != 0:
        added = False
        for i, cluster in enumerate(clusters):
            flag = True
            for feat in rules:
                if rules[feat] == 'exact':
                    if cluster['features'][feat] != row[feat]:
                        flag = False
                        break
            if flag:
                cluster['nodes'].append(row_id)
                cluster_id = i
                added = True
        if not added:
            cluster_id = len(clusters)
            clusters.append({'id': len(clusters), 'nodes': [row_id], 'features': row})
    else:
        cluster_id = 0
        clusters.append({'id': 0, 'nodes': [row_id], 'features': row})
    return cluster_id

def cluster_data(df, features, rules):
    df1 = df[features]
    df1 = df1.dropna()
    df1 = df1.reset_index(drop=True)
    df1['cluster'] = 0
    for i, row in df1.iterrows():
        if len(clusters) != 0:
            added = False
            for cluster in clusters:
                flag = True
                for feat in rules:
                    if rules[feat] == 'exact':
                        if cluster['features'][feat] != row[feat]:
                            flag = False
                            break
                if flag:
                    df1.loc[i, 'cluster'] = cluster['id']
                    cluster['nodes'].append(i)
                    added = True
            if not added:
                df1.loc[i, 'cluster'] = len(clusters)
                clusters.append({'id': len(clusters), 'nodes': [i], 'features': row})
        else:
            clusters.append({'id': 0, 'nodes': [i], 'features': row})
            df1.loc[i, 'cluster'] = 0
        print(i, len(clusters), df1.loc[i, 'cluster'])
        # break
    return df1

#TODO: Write logic to process stream data. (currently test data is printed)
from kafka import KafkaConsumer
from json import loads
from time import sleep
consumer = KafkaConsumer(
    'topic_test',
    bootstrap_servers=['kafka:9092'],
    auto_offset_reset='earliest',
    enable_auto_commit=True,
    group_id='my-group-id',
    value_deserializer=lambda x: loads(x.decode('utf-8'))
)
for event in consumer:
    event_data = event.value
    # Do whatever you want
    print(event_data)
    sleep(1)

# df1 = cluster_data(user_details_df, features, rules)

# df1.to_csv('clustered.csv')

