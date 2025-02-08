# Different import styles
import math
from random import randint
from datetime import datetime as dt
import os as operating_system

# Using math module
print(f"Square root of 16: {math.sqrt(16)}")
print(f"Pi: {math.pi}")

# Using random
print(f"\nRandom number between 1 and 10: {randint(1, 10)}")

# Using datetime
print(f"\nCurrent time: {dt.now()}")

# Using os
print(f"\nCurrent directory: {operating_system.getcwd()}")

# List all items in current directory
print("\nDirectory contents:")
for item in operating_system.listdir():
    print(f"- {item}") 