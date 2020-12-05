import sys


def f():
    if len(sys.argv) > 1:
        print("Cannot enter arguments.")
    return
    print("Am I dead?")


f()
