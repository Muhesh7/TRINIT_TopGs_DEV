from kafka import KafkaConsumer
from json import loads
from gen.app_auth_pb2 import AppAuthRequest
from gen.app_auth_pb2_grpc import AuthServiceStub
import grpc

def auth_rpc(id, secret):
    with grpc.insecure_channel('localhost:3001') as channel:
        stub = AuthServiceStub(channel)
        response = stub.AuthRPC(AppAuthRequest(app_id=id,app_secret=secret))
        return response

topic = 'topic_test'
group_id = 'app_group'
server = 'localhost:9092'

consumer =  KafkaConsumer(
                    topic,
                    bootstrap_servers=[server],
                    auto_offset_reset='earliest',
                    enable_auto_commit=True,
                    group_id=group_id,
                    value_deserializer=lambda x: loads(x.decode('utf-8'))
                )

check = {}

def start_consumption():

    for event in consumer:
        event_data = event.value
        rules = {}
        print(event_data)
        if event_data['app_id'] in check:
            rules = check[event_data['app_id']]
        else:
            print(event_data['app_id'],event_data['secret'])
            res=auth_rpc(id=event_data['app_id'],secret=event_data['secret'])
            print(res)


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

