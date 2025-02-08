# Basic decorator
def timer(func):
    from time import time
    def wrapper(*args, **kwargs):
        start = time()
        result = func(*args, **kwargs)
        end = time()
        print(f"{func.__name__} took {end - start:.4f} seconds")
        return result
    return wrapper

# Decorator with arguments
def repeat(times):
    def decorator(func):
        def wrapper(*args, **kwargs):
            for _ in range(times):
                result = func(*args, **kwargs)
            return result
        return wrapper
    return decorator

# Using decorators
@timer
def slow_function():
    import time
    time.sleep(1)
    print("Function executed")

@repeat(3)
def greet(name):
    print(f"Hello, {name}!")

# Test the decorators
print("Testing timer decorator:")
slow_function()

print("\nTesting repeat decorator:")
greet("Alice") 