# Basic error handling
try:
    number = int(input("Enter a number: "))
    result = 10 / number
    print(f"10 divided by {number} is {result}")
except ValueError:
    print("Please enter a valid number!")
except ZeroDivisionError:
    print("Cannot divide by zero!")
else:
    print("Calculation successful!")
finally:
    print("This always runs\n")

# Handling multiple errors with a single except
def get_item(lst, index):
    try:
        return lst[index]
    except (IndexError, TypeError) as e:
        return f"Error: {str(e)}"

print(get_item([1, 2, 3], 5))  # Index error
print(get_item(None, 0))      # Type error 