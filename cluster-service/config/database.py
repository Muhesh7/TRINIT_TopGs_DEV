from neo4j import GraphDatabase

class Database:
    def __init__(self, uri, user, password):
        self.driver = GraphDatabase.driver(uri, auth=(user, password))

    def close(self):
        self.driver.close()

    def get_rules_by_appId(self, app_id):
        with self.driver.session() as session:
            return session.execute_write(self._get_rules_by_appId, app_id)
        
    def create_app(self, app_id):
        with self.driver.session() as session:
            return session.execute_write(self._create_app, app_id)

    def create_rules_by_appId(self, app_id,rules):
        with self.driver.session() as session:
            return session.execute_write(self._create_rules_by_appId, app_id,rules)

    def create_data(self, app_id,data):
        with self.driver.session() as session:
            return session.execute_write(self._create_data, app_id,data)
        
    @staticmethod
    def _create_data(tx,app_id,datas:dict):
        json= "{"
        size = len(datas)
        i=0
        for data in datas:
            i+=1
            json+=str(data)+" : "+"'"+str(datas[data])+"'"
            if i<size:
                json+=", "
        json+="}"

        res0 =tx.run(
            "MERGE (n:Data"+json+")"
            "Return n")

        res = tx.run(
            "MATCH (n1:App)<-[:RULE_OFF]-(OtherNodes)"
            "WHERE n1.appname = {app_id}"
            "RETURN n1, OtherNodes", app_id=app_id
        )
        
        values = []
 
        for record in res:
            prop = record.values()[1]._properties
            values.append({'rule': prop['rule'], 'match': prop['match'], 'app_id': prop['app_id']})  

        for val in values:
            rule:str = val['rule']
            match:str = val['match']
            if match == 'exact':
                res = tx.run("MATCH (a:Data),(d:Data "+json+") "
                    "WHERE a."+rule+" = '"+datas[rule]+"' AND a.app_id = {app_id} AND NOT (a)-[:MATCH]-(d) AND NOT ID(a) =  ID(d)"
                    "MERGE (a)-[r:MATCH {type: 'exact', rule: {rule}}]-(d) "
                    "RETURN r,ID(a),ID(d)",app_id=app_id,rule=rule)
            elif match == 'partial':
                res = tx.run("MATCH (a:Data),(d:Data "+json+") "
                    "WHERE (a."+rule+" CONTAINS '"+datas[rule]+"' OR '"+datas[rule]+"' CONTAINS a."+rule+") "
                    "AND a.app_id = {app_id} AND NOT (a)-[:MATCH]-(d) AND NOT ID(a) = ID(d) "
                    "MERGE (a)-[r:MATCH {type: 'partial', rule: {rule}}]-(d) "
                    "RETURN r,ID(a),ID(d)",app_id=app_id,rule=rule)

    @staticmethod
    def _get_rules_by_appId(tx,app_id):
        res = tx.run(
            "MATCH (n1:App)<-[:RULE_OFF]-(OtherNodes)"
            "WHERE n1.appname = {app_id}"
            "RETURN n1, OtherNodes", app_id=app_id
        )
        values = []
        for record in res:
            prop = record.values()[1]._properties
            values.append({'rule': prop['rule'], 'match': prop['match'], 'app_id': prop['app_id']})  
        return values

    @staticmethod
    def _create_app(tx,app_id):
        res = tx.run("MERGE (n:App {appname: $app_id})"
                     "Return n",app_id=app_id)
        record = res.single()
        val = record.value()
        return val._properties['appname']

    @staticmethod
    def _create_rules_by_appId(tx,app_id,rules):
        values = []
        for rule in rules:
            match_type = rule['match_type']
            parameter = rule['parameter']
            tx.run(
            "MERGE (n:Rule {rule: $rule, match: $match, app_id: $app_id})",
            rule=parameter, match=match_type,  app_id=app_id
            )    
            res = tx.run(
                "MATCH (a:App {appname: $app_id}),(r:Rule {rule: $rule, match: $match, app_id: $app_id})"
                "WHERE NOT (a)<-[:RULE_OFF]-(r)"
                "CREATE (a)<-[:RULE_OFF]-(r)"
                "Return a,r",
                 app_id=app_id, rule= parameter, match= match_type
              
            )
            values.append(res)

        return values

connection = Database("bolt://localhost:7687", "neo4j", "password")

