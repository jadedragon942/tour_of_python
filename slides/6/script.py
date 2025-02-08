# Simple function with no parameters
def say_hello():
    print("Hello, World!")

# Function with parameters
def greet(name):
    return f"Hello, {name}!"

# Function with multiple parameters and default value
def calculate_total(price, tax_rate=0.1):
    return price + (price * tax_rate)

# Try the functions
say_hello()
print(greet("Alice"))
print(f"Total cost: ${calculate_total(100):.2f}")
print(f"Total cost with 20% tax: ${calculate_total(100, 0.2):.2f}") 