class Dog:
    def __init__(self, name, age):
        self.name = name
        self.age = age
    
    def bark(self):
        return f"{self.name} says Woof!"
    
    def info(self):
        return f"{self.name} is {self.age} years old"

# Create dog objects
buddy = Dog("Buddy", 5)
max = Dog("Max", 3)

# Use dog methods
print(buddy.bark())
print(max.info())

# Inheritance example
class Puppy(Dog):
    def __init__(self, name, age, training_level):
        super().__init__(name, age)
        self.training_level = training_level
    
    def bark(self):
        return f"{self.name} says Yip!"

charlie = Puppy("Charlie", 1, "Beginner")
print(charlie.bark())
print(charlie.info()) 