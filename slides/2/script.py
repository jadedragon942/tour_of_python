import random

def main():
    # starting a line with a '#' begins a comment. notice the next line does
    # not get executed.
    # print("hello world!")
    # meanwhile, random.randint(0, 9) in Python returns a random int in [0, 9].
    favorite_number = random.randint(0, 9)
    print("My favorite number is", favorite_number)

    # Let's explore Python types
    age = 25
    height = 1.75
    name = "Alice"
    is_student = True

    print(f"age is type: {type(age)}")
    print(f"height is type: {type(height)}")
    print(f"name is type: {type(name)}")
    print(f"is_student is type: {type(is_student)}")

if __name__ == "__main__":
    main()
