from kafka import KafkaConsumer, KafkaProducer
from json import loads, dumps
from json import dumps
from kafka import KafkaProducer

topic = 'app'
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
        
producer = KafkaProducer(
                bootstrap_servers=[server],
                value_serializer=lambda x: dumps(x).encode('utf-8')
        )

