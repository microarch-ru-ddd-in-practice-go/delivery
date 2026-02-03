import random
from dataclasses import dataclass

MIN_COORDINATE: int = 1
MAX_COORDINATE: int = 10


@dataclass(frozen=True)
class Location:
    x: int
    y: int


    # MagicMethod для валидации после инициализации объекта
    def __post_init__(self) -> None:
        if self.x < MIN_COORDINATE or self.x > MAX_COORDINATE:
            raise ValueError(f"Координата x имеет некорректное значение {self.x}")
        if self.y < MIN_COORDINATE or self.y > MAX_COORDINATE:
            raise ValueError(f"Координата y имеет некорректное значение {self.y}")

    # Так как используем @dataclass(frozen=True), то сеттеры и эквивалентность уже работают из коробки.
    # Если бы делали руками, то можно было бы реализовать следующим образом

    # def x(self) -> int:
    #     return self.x

    # def y(self) -> int:
    #     return self.y

    # def __eq__(self, other: "Location") -> bool:
    #     return self.x == other.x and self.y == other.y

    def distance_to(self, target: "Location") -> int:
        return abs(self.x - target.x) + abs(self.y - target.y)

    def __str__(self) -> str:
        return f"({self.x}, {self.y})"

    @staticmethod
    def new_random_location() -> "Location":
        random_x = random.randint(MIN_COORDINATE, MAX_COORDINATE)
        random_y = random.randint(MIN_COORDINATE, MAX_COORDINATE)
        return Location(x=random_x, y=random_y)
