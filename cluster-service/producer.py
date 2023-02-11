from time import sleep
from json import dumps
from kafka import KafkaProducer
import pandas as pd

user_details_df = pd.read_csv('user_details.csv', low_memory=False)

producer = KafkaProducer(
    bootstrap_servers=['kafka:9092'],
    value_serializer=lambda x: dumps(x).encode('utf-8')
)

for i, row in user_details_df.iterrows():
    data = {'id': i, 'name': row['name'], 'age': row['age'], 'country': row['country'], 'ip': row['ip']}
    producer.send('topic_test', value=data)
    sleep(0.5)
# for j in range(9999):
#     print("Iteration", j)
#     data = {'counter': j}
#     producer.send('topic_test', value=data)
#     sleep(0.5)