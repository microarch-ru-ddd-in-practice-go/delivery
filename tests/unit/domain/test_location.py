from dataclasses import FrozenInstanceError

import pytest

from core.domain.model.kernel.location import (
    MAX_COORDINATE,
    MIN_COORDINATE,
    Location,
)


class TestLocation:
    @pytest.mark.parametrize(
        "x, y, expected_error_msg",
        [
            (0, 1, "Координата x имеет некорректное значение 0"),
            (1, 11, "Координата y имеет некорректное значение 11"),
        ],
    )
    def test_create_with_invalid_coordinates_failed(
        self, x: int, y: int, expected_error_msg: str
    ) -> None:
        with pytest.raises(ValueError, match=expected_error_msg):
            Location(x=x, y=y)

    @pytest.mark.parametrize(
        "x, y",
        [
            (1, 1),
            (10, 10),
            (7, 3),
        ],
    )
    def test_create_with_valid_coordinates(self, x: int, y: int) -> None:
        location = Location(x=x, y=y)
        assert location.x == x
        assert location.y == y

    @pytest.mark.parametrize(
        "x1, y1, x2, y2, expected",
        [
            (1, 1, 1, 1, True),
            (1, 1, 1, 2, False),
            (1, 1, 2, 1, False),
            (5, 5, 5, 5, True),
            (10, 10, 10, 10, True),
            (5, 3, 3, 5, False),
        ],
    )
    def test_equality(self, x1: int, y1: int, x2: int, y2: int, expected: bool) -> None:
        loc1 = Location(x=x1, y=y1)
        loc2 = Location(x=x2, y=y2)
        assert (loc1 == loc2) == expected

    def test_modify_field_failed(self) -> None:
        location = Location(x=1, y=1)
        with pytest.raises(FrozenInstanceError):
            location.x = 10

    @pytest.mark.parametrize(
        "x1, y1, x2, y2, expected_distance",
        [
            (1, 1, 1, 1, 0),
            (1, 1, 1, 2, 1),
            (1, 1, 2, 1, 1),
            (1, 1, 5, 3, 6),
        ],
    )
    def test_distance_with_various_coordinates(
        self, x1: int, y1: int, x2: int, y2: int, expected_distance: int
    ) -> None:
        loc1 = Location(x=x1, y=y1)
        loc2 = Location(x=x2, y=y2)
        assert loc1.distance_to(loc2) == expected_distance

    def test_new_random_location_creates_valid_location(self) -> None:
        location = Location.new_random_location()
        assert isinstance(location, Location)
        assert MIN_COORDINATE <= location.x <= MAX_COORDINATE
        assert MIN_COORDINATE <= location.y <= MAX_COORDINATE
