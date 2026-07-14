"""A small, well-tested utility module — used as the 'good' validation commit.

Demonstrates: clear naming, docstrings, type hints, single-responsibility
functions, and explicit error handling — the kind of code that should pass
a Sonar quality gate cleanly.
"""

from __future__ import annotations


def celsius_to_fahrenheit(celsius: float) -> float:
    """Convert a Celsius temperature to Fahrenheit."""
    return (celsius * 9 / 5) + 32


def average(values: list[float]) -> float:
    """Return the arithmetic mean of a non-empty list of numbers.

    Raises:
        ValueError: if `values` is empty.
    """
    if not values:
        raise ValueError("Cannot compute the average of an empty list")
    return sum(values) / len(values)


def safe_divide(numerator: float, denominator: float) -> float | None:
    """Divide two numbers, returning None instead of raising on divide-by-zero."""
    if denominator == 0:
        return None
    return numerator / denominator
