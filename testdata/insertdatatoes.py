import random

from elasticsearch import Elasticsearch

es = Elasticsearch()

for i in range(0, 1000):
    arr = []
    for j in range(0, 4000):
        arr.append(random.random())
    es.index(index="features", doc_type="search_entry",  body={
        "id": i,
        "vector": arr
    })
    print i
