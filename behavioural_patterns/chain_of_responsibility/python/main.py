from typing import Dict
from dict_component import DictComponent
from numbers import Number


class NumericValueFinder(DictComponent):
    def handle(self, context: dict) -> dict:
        for k, v in list(context.items()):
            if not isinstance(v, Number):
                context.pop(k)
        return super().handle(context)


class ValueLimiter(DictComponent):
    def __init__(self, lower_bound: float, upper_bound: float) -> None:
        self.lower_bound = lower_bound
        self.upper_bound = upper_bound
        super().__init__()

    def handle(self, context: dict) -> dict:
        for k in list(context.keys()):
            v = context[k]
            if v < self.lower_bound or v >= self.upper_bound:
                context.pop(k)
        return super().handle(context)


class PowerRaiser(DictComponent):
    def __init__(self, exponent: int) -> None:
        self.exponent = exponent
        super().__init__()

    def handle(self, context: dict) -> dict:
        for k, v in list(context.items()):
            context[k] = v ** self.exponent
        return super().handle(context)


if __name__ == "__main__":
    c = {"a": 2, "b": "seven", "c": 23, "d": 42, "e": 76, "f": 124}
    nvf = NumericValueFinder()
    vl1 = ValueLimiter(lower_bound=1, upper_bound=100)
    pr = PowerRaiser(exponent=2)
    vl2 = ValueLimiter(lower_bound=1, upper_bound=1000)
    nvf.set_next(vl1)
    vl1.set_next(pr)
    pr.set_next(vl2)
    result = nvf.handle(c)
    print(f"result = {result}")
