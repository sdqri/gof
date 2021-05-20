from abstract_component import AbstractComponent


class DictComponent(AbstractComponent):

    next_component: "AbstractComponent" = None

    def handle(self, context: dict) -> dict:
        if self.next_component is not None:
            return self.next_component.handle(context)
        return context

    def set_next(self, next_component: "AbstractComponent"):
        self.next_component = next_component
