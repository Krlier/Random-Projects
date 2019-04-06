from calculator import Calculator

# Ask the user what operation to be done
def ask():
    while True:
        try:
            # Reads user input
            operation = input("What operation would you like to do? Please use this format: \"🐼 + 😎\": ")

            # Splits the string read into the operation array
            operation = operation.split(" ")

            # Use ord to get a number from the corresponding emoji
            emoji1 = ord(operation[0])
            emoji2 = ord(operation[1])

            break
        except Exception:
            print("Input in wrong format 🧐, please use this format: \"🐼 + 🐸\"")        

    return emoji1, emoji2, operation[1]

def main():
    # Ask() returns the emojis read and the operation
    emoji1, emoji2, operation = ask()

    # Create a new calculator
    c = Calculator()

    # Get the result operation from calculator
    result = c.calculate(emoji1, emoji2, operation)

    # Get the emoji corresponding to the number calculated
    newEmoji = chr(result)

    # Prints the emoji
    print(newEmoji)

# Starts main
main()