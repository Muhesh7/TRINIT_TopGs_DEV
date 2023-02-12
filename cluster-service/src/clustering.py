
features = ['name', 'age', 'country', 'ip']

rules = {
    'name': 'partial',
    'email': 'exact',
    'number': 'exact',
    'ip': 'exact',
}

clusters = []

def cluster_stream_data(row_id, row, features, rules):
    cluster_id = 0
    if len(clusters) != 0:
        added = False
        for i, cluster in enumerate(clusters):
            flag = True
            for feat in rules:
                if rules[feat] == 'exact':
                    if cluster['features'][feat] != row[feat]:
                        flag = False
                        break
            if flag:
                cluster['nodes'].append(row_id)
                cluster_id = i
                added = True
        if not added:
            cluster_id = len(clusters)
            clusters.append({'id': len(clusters), 'nodes': [row_id], 'features': row})
    else:
        cluster_id = 0
        clusters.append({'id': 0, 'nodes': [row_id], 'features': row})
    return cluster_id

def cluster_data(df, features, rules):
    df1 = df[features]
    df1 = df1.dropna()
    df1 = df1.reset_index(drop=True)
    df1['cluster'] = 0
    for i, row in df1.iterrows():
        if len(clusters) != 0:
            added = False
            for cluster in clusters:
                flag = True
                for feat in rules:
                    if rules[feat] == 'exact':
                        if cluster['features'][feat] != row[feat]:
                            flag = False
                            break
                if flag:
                    df1.loc[i, 'cluster'] = cluster['id']
                    cluster['nodes'].append(i)
                    added = True
            if not added:
                df1.loc[i, 'cluster'] = len(clusters)
                clusters.append({'id': len(clusters), 'nodes': [i], 'features': row})
        else:
            clusters.append({'id': 0, 'nodes': [i], 'features': row})
            df1.loc[i, 'cluster'] = 0
        print(i, len(clusters), df1.loc[i, 'cluster'])
        # break
    return df1

