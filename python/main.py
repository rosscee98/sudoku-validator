import numpy as np

input = [
    [
        [8,2,9,3,6,5,1,4,7],
        [6,4,3,7,1,8,5,1,9],
        [7,5,1,4,9,2,6,3,8],
        [3,1,8,5,2,7,4,9,6],
        [5,9,6,1,4,1,7,8,2],
        [4,7,2,9,8,6,3,1,5],
        [9,8,4,6,5,1,2,7,3],
        [2,6,7,8,3,4,9,5,1],
        [9,3,5,2,7,9,8,6,4]
    ],
    [
        [5,8,6,4,3,7,1,9,2],
        [1,9,4,5,8,2,3,6,7],
        [7,2,3,9,6,1,4,5,8],
        [2,4,7,1,9,8,6,3,5],
        [8,3,9,6,2,4,7,2,1],
        [6,5,1,7,2,3,8,4,9],
        [9,7,5,3,1,6,7,8,4],
        [3,1,8,2,4,5,9,7,6],
        [4,6,2,8,7,9,5,1,3]
    ],
    [
        [1,8,3,2,7,4,6,5,9],
        [9,7,4,5,8,6,3,2,1],
        [2,6,5,1,9,3,7,4,8],
        [5,9,2,8,3,1,4,6,7],
        [8,4,6,7,2,5,9,1,3],
        [7,3,1,4,6,9,2,8,5],
        [3,5,9,6,4,8,1,7,2],
        [6,1,7,3,5,2,8,9,4],
        [4,2,8,9,1,7,5,3,6]
    ]
]

validSequence = set([1,2,3,4,5,6,7,8,9])

def main():
    solutions = np.array(input)
    for solution in solutions:
        print(validate_solution(solution))

def validate_solution(solution):
    for i, row in enumerate(solution):
        if not (set(row) == validSequence):
            return f"Invalid - row {i+1}"

    for i, col in enumerate(solution.T):
        if not (set(col) == validSequence):
            return f"Invalid - col {i+1}"

    for i in range(0, 9, 3):
        for j in range(0, 9, 3):
            grid = solution[i:i+3, j:j+3]
            shapedGrid = np.reshape(grid, (9,))
            if not (set(shapedGrid) == validSequence):
                return f"Invalid - sub-grid {int(i / 3 + 1)}, {int(j / 3 + 1)}"

    return "Valid :)"

main()