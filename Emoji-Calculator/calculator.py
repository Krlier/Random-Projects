class Calculator():
    def __init__(self):
        pass

    def calculate(self, var1, var2, operation):
        if operation == '+':
            return (var1+var2)
        elif operation == '-':
            return (var1-var2)
        elif operation == '*':
            return (var1*var2)
        elif operation == '/':
            return (var1/var2)
        elif operation == '**':
            return (var1 ** var2)
        elif operation == '%':
            return (var1 % var2)
