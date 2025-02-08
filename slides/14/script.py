# Basic lambda function
square = lambda x: x**2
print(f"Square of 5: {square(5)}")

# Lambda with multiple arguments
sum_product = lambda a, b: (a + b, a * b)
print(f"\nSum and product of 3,4: {sum_product(3,4)}")

# Lambda with map
numbers = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x**2, numbers))
print(f"\nSquared numbers: {squared}")

# Lambda with filter
even_numbers = list(filter(lambda x: x % 2 == 0, numbers))
print(f"Even numbers: {even_numbers}")

# Lambda for sorting
people = [('Alice', 25), ('Bob', 30), ('Charlie', 20)]
# Sort by age
sorted_by_age = sorted(people, key=lambda x: x[1])
print(f"\nSorted by age: {sorted_by_age}")
# Sort by name
sorted_by_name = sorted(people, key=lambda x: x[0])
print(f"Sorted by name: {sorted_by_name}") 