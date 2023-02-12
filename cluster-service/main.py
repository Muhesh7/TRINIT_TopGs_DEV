from config.database import connection
from src.consumer import start_consumption


if __name__ == "__main__":
    # print(connection.create_app(app_id='app3'))
    # print(connection.create_rules_by_appId(app_id='app3',rules=[
    #     {'parameter': 'name', 'match_type': 'partial'},
    #     {'parameter': 'email', 'match_type': 'exact'},
    #     {'parameter': 'ip', 'match_type': 'exact'}
    # ]))
    # print(connection.get_rules_by_appId(app_id='app3'))
    start_consumption()
