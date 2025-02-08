# Basic list comprehension
numbers = [x for x in range(10)]
print("Numbers:", numbers)

# With condition
even_numbers = [x for x in range(10) if x % 2 == 0]
print("Even numbers:", even_numbers)

# With expression
squares = [x**2 for x in range(5)]
print("Squares:", squares)

# With strings
words = ["hello", "world", "python"]
upper_words = [word.upper() for word in words]
print("Uppercase words:", upper_words)

# Nested list comprehension
matrix = [[i+j for j in range(3)] for i in range(3)]
print("\nMatrix:")
for row in matrix:
    print(row) 