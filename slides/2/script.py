import random

def main():
    # starting a line with a '#' begins a comment. notice the next line does
    # not get executed.
    # print("hello world!")
    # meanwhile, random.randint(0, 9) in Python returns a random int in [0, 9].
    favorite_number = random.randint(0, 9)
    print("My favorite number is", favorite_number)

if __name__ == "__main__":
    main()
