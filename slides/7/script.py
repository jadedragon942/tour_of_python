# Create a dictionary
person = {
    "name": "Alice",
    "age": 25,
    "city": "New York"
}

# Access and print values
print("Name:", person["name"])
print("Age:", person.get("age"))

# Add a new key-value pair
person["job"] = "Developer"

# Update an existing value
person["age"] = 26

# Print all keys and values
print("\nKeys:", list(person.keys()))
print("Values:", list(person.values()))

# Loop through dictionary
print("\nPerson details:")
for key, value in person.items():
    print(f"{key}: {value}") 