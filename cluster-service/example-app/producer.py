import pandas as pd
from time import sleep
from kafka import KafkaProducer
from json import dumps
import os

topic = 'to_cluster'
server = 'localhost:9092'
                
        
producer = KafkaProducer(
                bootstrap_servers=[server],
                value_serializer=lambda x: dumps(x).encode('utf-8')
        )

cwd = os.getcwd()  

user_details_df = pd.read_csv('user.csv', low_memory=False)

for i, row in user_details_df.iterrows():
    data = {'app_id':'app3','secret':'secret',
            'name': row['name'], 'email': row['email'], 
            'ip': row['ip']
            }
    ans = producer.send(topic, value=data)
    print(data)
    sleep(0.1)