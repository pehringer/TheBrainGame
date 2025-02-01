# v v v Edit these settings v v v
filepath = "animation.json"
frame_ms = 500
x_units = 2
y_units = 2
z_units = 2
# ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^

import json
import matplotlib.pyplot as pyplot
import matplotlib.animation as animation
import mpl_toolkits.mplot3d.art3d as art3d

def plot_cube(ax, x, y, z):
    vertices = [
        [x, y, z], [x + 1, y, z], [x + 1, y + 1, z], [x, y + 1, z],
        [x, y, z + 1], [x + 1, y, z + 1], [x + 1, y + 1, z + 1], [x, y + 1, z + 1]
    ]
    faces = [
        [vertices[j] for j in [0, 1, 2, 3]], [vertices[j] for j in [4, 5, 6, 7]],
        [vertices[j] for j in [0, 1, 5, 4]], [vertices[j] for j in [2, 3, 7, 6]],
        [vertices[j] for j in [0, 3, 7, 4]], [vertices[j] for j in [1, 2, 6, 5]]
    ]
    ax.add_collection3d(art3d.Poly3DCollection(faces, color='grey', alpha=0.5, edgecolor='grey'))

def update(frame, data, ax):
	ax.cla()
	ax.set_axis_off()
	ax.set_xlim([0, x_units])
	ax.set_ylim([0, y_units])
	ax.set_zlim([0, z_units])
	points = data[frame] if frame < len(data) else []
	for p in points:
		plot_cube(ax, p["x"], p["y"], p["z"])
	return ax

fig = pyplot.figure()
ax = fig.add_subplot(111, projection='3d')
with open(filepath, 'r') as file:
	data = json.load(file)
ani = animation.FuncAnimation(fig, update, frames=len(data), fargs=(data, ax), interval=frame_ms, blit=False)
ani.save("animation.gif", writer="pillow")
pyplot.show()
