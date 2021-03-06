import sys
from typing import Union, List

INF = 2 * 10 ** 14


def check(X, n, hs, ss):
    data = [0] * n

    for h, s in zip(hs, ss):
        j = (X - h) // s
        if j < 0:
            return False
        data[min(j, n - 1)] += 1

    if data[0] >= 2:
        return False

    for i in range(1, n):
        data[i] += data[i - 1]
        if data[i] >= i + 2:
            return False

    return True


def main():
    n = read_num()

    hs = list()
    ss = list()

    for _ in range(n):
        IS = read_nums()
        hs.append(IS[0])
        ss.append(IS[1])

    l, u = 0, INF

    while l < u:
        m = (l + u) // 2
        if check(m, n, hs, ss):
            u = m
        else:
            l = m + 1
    ans = l

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


if __name__ == "__main__":
    main()
