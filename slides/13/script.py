# String methods
text = "  Hello, Python World!  "
print(f"Original: '{text}'")
print(f"Stripped: '{text.strip()}'")
print(f"Upper: '{text.upper()}'")
print(f"Lower: '{text.lower()}'")

# Splitting and joining
words = text.split()
print(f"\nSplit into words: {words}")
print(f"Joined back: '{' '.join(words)}'")

# Finding and replacing
sentence = "Python is amazing and Python is fun"
print(f"\nOriginal: {sentence}")
print(f"Replaced: {sentence.replace('Python', 'Coding')}")
print(f"First 'Python' at index: {sentence.find('Python')}")

# String formatting
name = "Alice"
age = 25
# Old style
print("\n%s is %d years old" % (name, age))
# Format method
print("{} is {} years old".format(name, age))
# f-string (modern)
print(f"{name} is {age} years old") 