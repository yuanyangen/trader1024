from jqdatasdk import *
auth('18513530687','yyg@J0inquant') #账号是申请时所填写的手机号；密码为聚宽官网登录密码

df = get_all_securities(types=["futures"], date=None)
for data in df.index:
    get_bars(data, 10000, unit='1d',
             fields=['date', 'open', 'close', 'high', 'low', 'volume', 'money','open_interest'],
             include_now=False, end_dt=None,df=True)
    print(data)
