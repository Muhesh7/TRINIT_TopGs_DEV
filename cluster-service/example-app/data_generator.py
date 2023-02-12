from faker import Faker
import numpy as np
from faker.providers import internet
import pandas as pd
import json
from random import randint

fake = Faker()
fake.add_provider(internet)

def get_fake_data():
    return {
        'name': fake.name(),
        'email': fake.email(),
        'mobile': "%0.12d" % randint(0,999999999999),
        'address': fake.address(),
        'ip': fake.ipv4_private()
    }

with open('user_details.json', 'w') as f:
    data = []
    for _ in range(10):
        data.append(get_fake_data())
    json.dump(data, f)

csv_file = open('user_details.csv', 'wb')

pd.read_json('user_details.json').to_csv('user_details.csv', index=False)
