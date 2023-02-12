import pandas as pd
from time import sleep
from kafka import KafkaProducer
from json import dumps
import os

topic = 'topic_test'
group_id = 'app_group'
server = 'localhost:9092'
                
        
producer = KafkaProducer(
                bootstrap_servers=[server],
                value_serializer=lambda x: dumps(x).encode('utf-8')
        )

cwd = os.getcwd()  

user_details_df = pd.read_csv('user_details.csv', low_memory=False)

for i, row in user_details_df.iterrows():
    data = {'app_id':'2','secret':'secret','id': i, 'name': row['name'], 'email': row['email'],'mobile': row['mobile'], 'address': row['address'], 'ip': row['ip']}
    ans = producer.send('topic_test', value=data)
    print(ans)
    sleep(0.1)