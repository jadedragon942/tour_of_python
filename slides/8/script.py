# Create sets
fruits = {"apple", "banana", "orange"}
citrus = {"lemon", "orange", "lime"}

# Add and remove items
fruits.add("grape")
fruits.discard("banana")

print("Fruits:", fruits)
print("Citrus:", citrus)

# Set operations
print("\nUnion:", fruits | citrus)
print("Intersection:", fruits & citrus)
print("Difference (fruits - citrus):", fruits - citrus)

# Remove duplicates from a list
numbers = [1, 2, 2, 3, 3, 3, 4, 4, 4, 4]
unique_numbers = list(set(numbers))
print("\nOriginal list:", numbers)
print("Unique numbers:", unique_numbers) 