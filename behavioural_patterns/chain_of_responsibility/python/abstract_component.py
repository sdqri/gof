from abc import ABC, abstractmethod


class AbstractComponent(ABC):
    @abstractmethod
    def handle(self, request):
        pass

    @abstractmethod
    def set_next(self, next_component: "AbstractComponent") -> "AbstractComponent":
        pass
