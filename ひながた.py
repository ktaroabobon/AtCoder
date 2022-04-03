def main():
    print('hello world')


def s2i(s: str) -> int:
    return int(s)


def i2s(i: int) -> str:
    return str(i)


def sReader():
    """
     文字列、１単語
     e.g.)
     foo
     """
    str_returned = input()

    return str_returned


def iReader():
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
