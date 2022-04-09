import sys
from typing import Union, List


def main():
    n = read_num()
    testimonies = [[] for _ in range(n)]
    ans = 0

    for i in range(n):
        num = read_num()
        for _ in range(num):
            t = read_nums()
            testimonies[i].append((t[0] - 1, t[1]))

    for i in range(1 << n):
        f = True
        for j in range(n):
            if not i >> j & 1:
                continue
            for p, d in testimonies[j]:
                if i >> p & 1 != d:
                    f = False
                    break
            if not f:
                break
        if f:
            ans = max(ans, bin(i)[2:].count('1'))

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
