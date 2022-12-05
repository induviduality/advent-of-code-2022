# Python code to plot a list of points
import matplotlib.pyplot as plt
import numpy as np

# Output from the Go program
coordinatesList = [[-1, -1], [42, 0], [42, 6], [42, 3], [46, 3], [46, 6], [46, 0], 
          [-1, -1], [9, 0], [9, 5], [10, 6], [12, 6], [13, 5], [13, 3], [10, 3],
          [13, 3], [13, 0], [-1, -1], [0, 0], [0, 6], [3, 3], [6, 6], [6, 0], 
          [-1, -1], [6, 5], [-1, -1], [39, 5], [38, 6], [36, 6], [35, 5], 
          [35, 1], [36, 0], [38, 0], [39, 1], [-1, -1], [39, 6], [-1, -1], 
          [20, 6], [18, 6], [16, 4], [16, 2], [18, 0], [20, 0], [20, 3],
          [19, 3], [-1, -1], [24, 3], [-1, -1], [24, 0], [23, 0], [25, 0], 
          [24, 0], [24, 6], [25, 6], [23, 6], [-1, -1], [28, 6], [-1, -1], 
          [28, 0], [31, 0], [32, 1], [32, 2], [31, 3], [29, 3], [28, 4], 
          [28, 5], [29, 6], [32, 6]]

# Group all the coordinates traversed without jumping
coordinatesGroup= []
for coordinates in coordinatesList:
    # removing the [-1, -1] markers
    if coordinates == [-1, -1]:
        coordinatesGroup.append([])
    else:
        coordinatesGroup[-1].append(coordinates)

# plot the points and connecting all of them
plt.figure(figsize=(25, 5))
for group in coordinatesGroup:
  x = [p[0] for p in group]
  y = [p[1] for p in group]
  plt.plot(x, y, "o-")
plt.show()
