def f():
    raise Exception()


try:
    f()
    print("Am I dead?")
except Exception:
    pass
