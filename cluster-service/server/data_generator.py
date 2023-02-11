from faker import Faker
import numpy as np
from faker.providers import internet
import pandas as pd
import json

fake = Faker()
fake.add_provider(internet)

def get_fake_data():
    return {
        'name': fake.name(),
        'email': fake.email(),
        'age': fake.random_int(min=18, max=100),
        'address': fake.address().split('\n')[0],
        'city': fake.city(),
        'state': fake.state(),
        'country': fake.country(),
        'zipcode': fake.zipcode(),
        'gender': np.random.choice(["M", "F"], p=[0.5, 0.5]),
        'ip': fake.ipv4_private()
    }

with open('user_details.json', 'w') as f:
    data = []
    for _ in range(1000):
        data.append(get_fake_data())
    json.dump(data, f)

csv_file = open('user_details.csv', 'wb')

pd.read_json('user_details.json').to_csv('user_details.csv', index=False)
