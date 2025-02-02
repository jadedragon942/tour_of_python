import random

def main():
    # rand.Intn(10) in Go returns a random int in [0, 10).
    # random.randint(0, 9) in Python returns a random int in [0, 9].
    favorite_number = random.randint(0, 9)
    print("My favorite number is", favorite_number)

if __name__ == "__main__":
    main()
