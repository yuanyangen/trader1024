import pandas as pd
import requests
from urllib.parse import urlencode

EastmoneyKlines = {
    'f51': 'Date',
    'f52': 'Open',                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   n                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   
    'f53': 'Close',
    'f54': 'High',
    'f55': 'Low',
    # 'f56': 'Volume',
    # 'f57': '成交额',
    # 'f58': '振幅',
    # 'f59': '涨跌幅',
    # 'f60': '涨跌额',
    # 'f61': '换手率',
}

EastmoneyHeaders = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 6.3; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko',
    'Accept': '*/*',
    'Accept-Language': 'zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2',
    'Referer': 'http://quote.eastmoney.com/center/gridlist.html',
}


def get_futures_base_info() -> pd.DataFrame:
    """
    获取四个期货交易所全部期货基本信息
    """
    params = (
        ('np', '1'),
        ('fltt', '2'),
        ('invt', '2'),
        ('fields', 'f1,f2,f3,f4,f12,f13,f14'),
        ('pn', '1'),
        ('pz', '300000'),
        ('fid', 'f3'),
        ('po', '1'),
        ('fs', 'm:113,m:114,m:115,m:8'),
        ('forcect', '1'),
    )
    rows = []
    cfg = {
        113: '上期所',
        114: '大商所',
        115: '郑商所',
        8: '中金所'
    }
    response = requests.get(
        'https://push2.eastmoney.com/api/qt/clist/get', headers=EastmoneyHeaders, params=params)
    for item in response.json()['data']['diff']:
        code = item['f12']
        name = item['f14']
        secid = str(item['f13']) + '.' + code
        belong = cfg[item['f13']]
        row = [code, name, secid, belong]
        rows.append(row)
    columns = ['期货代码', '期货名称', 'secid', '归属交易所']
    return pd.DataFrame(rows, columns=columns)


def get_ftures_k_history(secid: str, beg: str = '19000101', end: str = '20500101', klt: int = 101,
                         fqt: int = 2) -> pd.DataFrame:
    """
    获取k线数据

    Parameters
    ----------
    secid : 根据 get_futures_base_info 函数获取
    获取4个交易所期货数据，取 secid 列来获取 secid
    beg : 开始日期 例如 20200101
    end : 结束日期 例如 20200201
    klt : k线间距 默认为 101 即日k
            klt : 1 1 分钟
            klt : 5 5 分钟
            klt : 101 日
            klt : 102 周
    fqt: 复权方式
            不复权 : 0
            前复权 : 1
            后复权 : 2

    Return
    ------
    DateFrame : 包含期货k线数据
    """

    fields = list(EastmoneyKlines.keys())
    columns = list(EastmoneyKlines.values())
    fields2 = ",".join(fields)

    params = (
        ('fields1', 'f1,f2,f3,f4,f5,f6,f7,f8,f9,f10,f11,f12,f13'),
        ('fields2', fields2),
        ('beg', beg),
        ('end', end),
        ('rtntype', '6'),
        ('secid', secid),
        ('klt', f'{klt}'),
        ('fqt', f'{fqt}'),
    )
    base_url = 'https://push2his.eastmoney.com/api/qt/stock/kline/get'
    url = base_url + '?' + urlencode(params)
    json_response = requests.get(url, headers=EastmoneyHeaders).json()

    data = json_response['data']
    if data is None:
        print(secid, '无数据')
        return None
    # code = data['code']
    # name
    # name = data['name']
    klines = data['klines']

    rows = []
    for _kline in klines:
        kline = _kline.split(',')
        rows.append(kline)

    df = pd.DataFrame(rows, columns=columns)

    return df


if "__main__" == __name__:
    # 获取全市场期货基本数据，即下面的 4 列信息
    # ['期货代码', '期货名称', 'secid', '归属交易所']
    futures_info_df = get_futures_base_info()
    futures_info_df.to_csv('期货信息表.csv', index=None, encoding='utf-8-sig')
    print('期货基本信息表如下')
    secids = futures_info_df['secid']
    futures_names = futures_info_df['期货名称']
    # 随便选了 1 个期货来测试
    for i in range(0, len(secids)):
        futures_name = futures_names[i]
        if '主力' not in futures_name:
            continue
        # 随便选的期货名称
        secid = secids[i]
        print(f'正在获取期货：{futures_name} 的历史 k 线数据')
        # 获取期货历史 日k 线数据
        # secid 可以在文件 期货信息表.csv 中查看，更改 secid 即可获取不同的期货数据
        df = get_ftures_k_history(secid)
        df.to_csv(f'daily/{futures_name}.csv', index=None, encoding='utf-8-sig')
        print(f'期货 : {futures_name} 的 k 线数据已保存到文件 {futures_name}.csv 中')
