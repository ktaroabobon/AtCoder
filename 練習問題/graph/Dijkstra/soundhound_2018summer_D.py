"""
問題URL:
"""

import math
import sys
from collections import deque
from typing import Union, List

INF = 2 * 10 ** 14
CONST = 998244353

global g
global yen_dist
global snk_dist
global q


class Edge(object):
    def __init__(self, to, yen, snk):
        self.to = to
        self.yen = yen
        self.snk = snk

    def __lt__(self, other):
        if isinstance(other, Edge):
            return
        return self.to < other.to


def bfs(sp, g, dist, q: deque, yen=True):
    dist[sp] = 0
    q.append(sp)

    while len(q) > 0:
        v = q.popleft()

        for nv in g[v]:
            if yen:
                if dist[nv.to] > dist[v] + nv.yen:
                    dist[nv.to] = dist[v] + nv.yen
                    q.append(nv.to)
            else:
                if dist[nv.to] > dist[v] + nv.snk:
                    dist[nv.to] = dist[v] + nv.snk
                    q.append(nv.to)

    return dist


def main():
    N, M, s, t = read_nums()
    g = [[] for _ in range(N)]

    for _ in range(M):
        u, v, a, b = read_nums()
        g[u - 1].append(Edge(v - 1, a, b))
        g[v - 1].append(Edge(u - 1, a, b))

    yen_dist = [INF] * N
    snk_dist = [INF] * N

    q = deque()
    yen_dist = bfs(s - 1, g, yen_dist, q)

    q = deque()
    snk_dist = bfs(t - 1, g, snk_dist, q, yen=False)

    total = list()

    for y, s in zip(yen_dist, snk_dist):
        total.append(y + s)

    ans = []
    tmp = INF
    init = int(1e15)
    for t in reversed(total):
        tmp = min(tmp, t)
        ans.append(init - tmp)

    for a in ans.__reversed__():
        print(a)


def split_without_empty(strs: str) -> List[str]:
    """
    文字列を分割してlistに格納し返す
    Args:
        strs: 複数の文字

    Returns: listに複数の文字列を格納されたもの
    Examples: foo boo -> [foo, boo]
    """
    return strs.split(' ')


def split2int(strs: List[str]) -> List[int]:
    """
    文字列型のlistを数値型のlistに変換する

    Args:
        strs: 数値が文字列型のlist

    Returns: 数値型のlist
    Examples: ['100', '200'] -> [100, 200]
    """
    return [int(n) for n in strs]


def split2str(ints: List[int]) -> List[str]:
    """
    数値型のlistを文字列型のlistに変換する

    Args:
        ints: 数値型のlist

    Returns: 文字列型のlist
    Examples: [100, 200] -> ['100', '200']
    """
    return [str(n) for n in ints]


def s2i(s: str) -> int:
    return int(s)


def i2s(i: int) -> str:
    return str(i)


def b2i(b: bool) -> int:
    return int(b)


def i2b(i: int) -> bool:
    return bool(i)


def read_str() -> str:
    """
     文字列、１単語
     e.g.)
     foo
     """

    return sys.stdin.readline().rstrip()


def read_strs() -> List[str]:
    """
    文字列、複数単語

    Returns: List[str]
    Examples:
        foo, boo
    """

    return split_without_empty(read_str())


def read_num() -> Union[int, float]:
    """
    数値
    e.g.)
    10
    """

    return s2i(read_str())


def read_nums() -> Union[List[int], List[float]]:
    """
    数値
    e.g.)
    [10, 20]
    """
    return split2int(read_strs())


def aCb(a, b: int) -> int:
    """
    二項定理

    Args:
        a (int)
        b (int)

    Returns:
        二項定理の値
    """
    r = 1
    if a < b * 2:
        b = a - b

    for i in range(b):
        r *= (a - i + 1) / i

    return r


def get_distance(p1, p2: List[int]) -> Union[int, float]:
    """
    2点間距離

    Args:
        p1(List[int]): 座標
        p2(List[int]): 座標

    Returns:
        距離
    """
    d = 0
    for x1, x2 in zip(p1, p2):
        d += (x1 - x2) ** 2
    return math.sqrt(d)


if __name__ == "__main__":
    main()
