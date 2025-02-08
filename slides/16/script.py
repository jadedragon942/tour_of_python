# Simple generator function
def count_up_to(n):
    i = 1
    while i <= n:
        yield i
        i += 1

# Fibonacci generator
def fibonacci():
    a, b = 0, 1
    while True:
        yield a
        a, b = b, a + b

# Using generators
print("Counting up to 5:")
for num in count_up_to(5):
    print(num)

print("\nFirst 8 Fibonacci numbers:")
fib = fibonacci()
for _ in range(8):
    print(next(fib))

# Generator expression
squares = (x**2 for x in range(5))
print("\nSquares using generator expression:")
for square in squares:
    print(square)

# Memory comparison
import sys
list_comp = [x for x in range(1000)]
gen_exp = (x for x in range(1000))
print(f"\nList comprehension size: {sys.getsizeof(list_comp)} bytes")
print(f"Generator expression size: {sys.getsizeof(gen_exp)} bytes") 