from src.consumer import start_consumption


def start_kafka():
    res = "success"
    try:
        start_consumption()
    except Exception as e:
        print(e)
        res = e
    print(res)

if __name__ == "__main__":
    start_kafka()
