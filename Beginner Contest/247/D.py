"""
問題URL: https://atcoder.jp/contests/abc247/tasks/abc247_d
"""

import bisect
import math
import sys
from typing import Union, List

INF = 2 * 10 ** 14


def main():
    n = read_num()
    balls = [0]
    sum_values = [0]
    total_balls = 0
    per_ans = 0

    for _ in range(n):
        x = read_nums()
        if x[0] == 1:
            if len(balls) == 0:
                balls.append(x[2])
                sum_values.append(x[1] * x[2])
            else:
                balls.append(balls[-1] + x[2])
                sum_values.append(sum_values[-1] + x[1] * x[2])
        else:
            total_balls += x[1]
            idx = bisect.bisect_left(balls, total_balls)
            if balls[idx] == x[1]:
                ans = sum_values[idx] - per_ans
                per_ans += ans
            else:
                ans = sum_values[idx] - per_ans - (
                            (sum_values[idx] - sum_values[idx - 1]) // (balls[idx] - balls[idx - 1])) * (
                                  balls[idx] - total_balls)
                per_ans += ans
            # per_total_balls = total_balls

            print(ans)


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
