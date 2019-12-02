from requests_html import HTMLSession
from time import sleep
from fake_useragent import UserAgent
import json

# 1.生成session
session = HTMLSession()

# 2. 发起第一次请求，获取cookie
a = session.get(
    url='https://www.lagou.com/jobs/list_go?labelWords=&fromSearch=true&suginput=',
    headers={
        'useragent': UserAgent().random})
sleep(5)

# 3.构建请求
url = 'https://www.lagou.com/jobs/positionAjax.json'
ua = UserAgent()
headers = {
    'user-agent': ua.random,
    'referer': 'https://www.lagou.com/jobs/list_go?city=%E6%B7%B1%E5%9C%B3&cl=false&fromSearch=true&labelWords=&suginput=&labelWords=hot',
    'Origin': 'https://www.lagou.com',
    'Accept': 'application/json, text/javascript, */*; q=0.01',
    'Content-Type': 'application/x-www-form-urlencoded; charset=UTF-8'}

data = {
    'first': 'true',
    'pn': '1',
    'kd': 'go',
    'city': '深圳',
    'needAddtionalResult': 'false'
}

# 4.发起请求
r = session.post(
    url=url,
    headers=headers,
    data=data
)
content = r.content

# 5.保存数据
json_msg = json.loads(content)

print(type(json_msg))
with open('lagou.json', 'w') as fd:
    json.dump(json_msg, fd)