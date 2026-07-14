"""Intentionally poor-quality module — used as the 'bad' validation commit.

Deliberately includes the kinds of issues SonarQube is expected to catch:
- unused import (code smell)
- bare `except:` (code smell / reliability)
- deep nesting / high cyclomatic complexity (maintainability)
- mutable default argument (bug-prone pattern)
- a hardcoded, fake-looking secret pattern (to trigger a Security Hotspot review)
- no accompanying unit tests, so it drags down overall coverage

NOTE: the "API_KEY" below is a placeholder string, not a real credential —
it exists purely so the scanner flags a Security Hotspot for review.
"""

import os
import sys
import json  # unused import on purpose

API_KEY = "FAKE-PLACEHOLDER-KEY-1234567890"  # intentional hotspot trigger


def process(data, cache={}):  # mutable default argument
    if data:
        if isinstance(data, list):
            for item in data:
                if item:
                    if isinstance(item, dict):
                        for key in item:
                            if key == "value":
                                if item[key] > 0:
                                    cache[key] = item[key]
                                else:
                                    cache[key] = 0
    return cache


def risky_operation(value):
    try:
        return 100 / value
    except:  # bare except - swallows everything, including KeyboardInterrupt
        return None


def unused_variable_example():
    unused = "this variable is never used"
    result = 1 + 1
    return result
