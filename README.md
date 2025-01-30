# VidFlex

A Q-Learning-based adaptive video streaming optimizer, entirely written in Go, designed to maximize user satisfaction while minimizing bandwidth costs. The current implementation operates on synthetic data with simplified states and utilizes Q-Learning with an epsilon-greedy policy.

![VidFlex Banner](/img/vidflex.png)

## Features

- [x] Simulate concurrent users watching the same video stream with randomized network bandwidth ranges (0–1 Mbps, 1–3 Mbps, 3–5 Mbps, 5+ Mbps) and device types (mobile, tablet, desktop/TV).
- [x] Adjust video quality by modifying one metric at a time (resolution, bitrate, frame rate) through discrete actions.
- [x] Update the policy using a reward function that combines engagement score (prioritizing higher quality), bandwidth cost penalties, and a collaborative adjustment based on the average satisfaction of other concurrent users in similar states.  
       Formula:  
       $$Q(s, a) \leftarrow Q(s, a) + \alpha \left[\text{reward} + \gamma \max_{a'} Q(s', a') + 0.2 \times \text{average satisfaction}\right]$$

- [x] Implement an epsilon-greedy policy with an initial epsilon = 0.8 and epsilon decay of 5% per 100 episodes.
- [x] **System assumptions**: Concurrent users are managed by a single agent via `goroutines`; simulated bandwidth dips reduce the bandwidth tier by one level with 10% probability.
- [x] Train the agent and benchmark performance against a static policy (fixed 720p, 1 Mbps, 30 fps) via command line (see the below section for more details).

## Benchmark Results

The performance of the RL model after 10,000 episodes of training on a 15-inch MacBook Air with an M3 Chip (8-Core CPU, 10-Core GPU, and 16GB Unified Memory) is shown below.

![Performance Benchmark](/img/benchmark.png)

## Getting Started

1. Ensure Go is installed (this project was developed with Go 1.23.4). Start by cloning the repository:

   ```bash
   git clone https://github.com/oadultradeepfield/vidflex.git
   cd vidflex
   ```

2. To run the training, execute this command with customizable parameters:

   ```bash
   go run cmd/train/main.go \
   -episodes 2000 \
   -users 500 \
   -alpha 0.15 \
   -gamma 0.95 \
   -epsilon 0.2 \
   -decay-every 200 \
   -decay-rate 0.05
   ```

3. To generate benchmark results, use the preconfigured script files:

   ```bash
   chmod +x benchmark.sh
   ./benchmark.sh
   ```

   The benchmark results will be saved in the `results.csv` file located in the root directory.

## Future Roadmap

- [ ] Implement advanced algorithms like Deep Q-Networks (DQN) to stabilize learning as user numbers scale, using techniques like experience replay and target networks.
- [ ] Integrate alternative exploration strategies (e.g., Upper Confidence Bound [UCB], Thompson Sampling) to improve state-action space exploration beyond epsilon-greedy.
- [ ] Expand state definitions with real-world network parameters (e.g., latency, jitter, packet loss) and device-specific capabilities (e.g., codec support, screen resolution).

## License

This project is licensed under the GNU General Public License v3.0 . See the [`LICENSE`](/LICENSE) file for details.
