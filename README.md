# TheBrainGame
A 3D cellular automaton where each cell exists in one of two states: resting (0) or spiking (1).
Each cell has a unique threshold (0-7) and six unique connections, one for each adjacent neighbor.
Connections can be either excitatory (+1) or inhibitory (-1).
A cell's state is determined by computing the dot product of its connections and their corresponding neighbor states.
If the result meets or exceeds the cell’s threshold, it spikes (1). Otherwise, it rests (0).
![Current Progress](output.gif)
