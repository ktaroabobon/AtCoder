import sys
from typing import Union, List


def main():
    num = read_num()

    r = 1
    f = 0

    for i in range(1, num + 1):
        x = i
        cnt = 0
        while True:
            if x % 2 == 1:
                break
            x /= 2
            cnt += 1

        if f < cnt:
            f = cnt
            r = i

    print(r)


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
    str_returned = sys.stdin.readline().rstrip()

    return str_returned


def read_strs() -> List[str]:
    """
    文字列、複数単語

    Returns: List[str]
    Examples:
        foo, boo
    """
    strs = read_str()

    return split_without_empty(strs)


def read_num() -> Union[int, float]:
    """
    数値
    e.g.)
    10
    """
    s = read_str()

    num_returned = s2i(s)

    return num_returned


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
