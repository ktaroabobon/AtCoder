import sys


def main():
    print('hello world')
    print(sys.version)


def s2i(s: str) -> int:
    return int(s)


def i2s(i: int) -> str:
    return str(i)


def b2i(b: bool) -> int:
    return int(b)


def i2b(i: int) -> bool:
    return bool(i)


def str_reader():
    """
     文字列、１単語
     e.g.)
     foo
     """
    str_returned = sys.stdin.readline().rstrip()

    return str_returned


def int_reader():
    """
    数値、１整数
    e.g.)
    10
    """
    str = sReader()

    num_returned = s2i(str)

    return num_returned


if __name__ == "__main__":
    main()
