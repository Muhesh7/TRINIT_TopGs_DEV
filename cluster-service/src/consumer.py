from json import loads
from gen.app_auth_pb2 import AppAuthRequest
from gen.app_auth_pb2_grpc import AuthServiceStub
import grpc
from config.database import connection
from kafka import KafkaConsumer

def auth_rpc(id, secret):
    with grpc.insecure_channel('localhost:3001') as channel:
        stub = AuthServiceStub(channel)
        response = stub.AuthRPC(AppAuthRequest(app_id=id,app_secret=secret))
        return response

topic='to_cluster'
group_id='app_group'
server='localhost:9092'

def start_consumption():
    consumer = KafkaConsumer(
                    topic,
                    group_id=group_id,
                    auto_offset_reset='earliest',
                    enable_auto_commit=True,
                    bootstrap_servers=[server],
                    value_deserializer=lambda x: loads(x.decode('utf-8'))
                )
    print(consumer.config)
    print(consumer.bootstrap_connected())
    for event in consumer:
        connection.create_data(app_id=event.value['app_id'],data=event.value)

