
import ast

program = """
import ast
my_tree = ast.parse("3 + 4*x")
print(ast.dump(my_tree))

def pi():
    return 3.1415
print(2 * pi())

x = 1
def print_x():
    print(x)
    if False: x = 0
print_x()
"""
tree = ast.parse(program)
print(ast.dump(tree))
