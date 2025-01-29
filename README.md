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

**Device**: 15-inch MacBook Air with M3 Chip (8-Core CPU, 10-Core GPU, and 16GB Unified Memory)

### Users: 1

| Episodes | Time   | Dynamic Reward | Static Reward | Dynamic Penalty | Static Penalty |
| -------- | ------ | -------------- | ------------- | --------------- | -------------- |
| 500      | 6.81ms | 0.4920         | 0.3420        | 0.5000          | 0.5000         |
| 1000     | 4.85ms | 0.4320         | 0.3420        | 0.3000          | 0.5000         |
| 2000     | 9.79ms | 0.4320         | 0.2394        | 0.3000          | 0.5000         |

### Users: 100

| Episodes | Time     | Dynamic Reward | Static Reward | Dynamic Penalty | Static Penalty |
| -------- | -------- | -------------- | ------------- | --------------- | -------------- |
| 500      | 109.41ms | 0.4202         | 0.3276        | 0.3090          | 0.5000         |
| 1000     | 218.07ms | 0.4216         | 0.3256        | 0.3000          | 0.5000         |
| 2000     | 445.60ms | 0.4190         | 0.3297        | 0.3000          | 0.5000         |

### Users: 500

| Episodes | Time     | Dynamic Reward | Static Reward | Dynamic Penalty | Static Penalty |
| -------- | -------- | -------------- | ------------- | --------------- | -------------- |
| 500      | 667.71ms | 0.4148         | 0.3315        | 0.3162          | 0.5000         |
| 1000     | 1.40s    | 0.4177         | 0.3322        | 0.3100          | 0.5000         |
| 2000     | 2.87s    | 0.4177         | 0.3340        | 0.3040          | 0.5000         |

### Users: 1000

| Episodes | Time  | Dynamic Reward | Static Reward | Dynamic Penalty | Static Penalty |
| -------- | ----- | -------------- | ------------- | --------------- | -------------- |
| 500      | 1.47s | 0.4148         | 0.3325        | 0.3156          | 0.5000         |
| 1000     | 3.07s | 0.4171         | 0.3304        | 0.3081          | 0.5000         |
| 2000     | 6.29s | 0.4180         | 0.3325        | 0.3033          | 0.5000         |

### Users: 2000

| Episodes | Time   | Dynamic Reward | Static Reward | Dynamic Penalty | Static Penalty |
| -------- | ------ | -------------- | ------------- | --------------- | -------------- |
| 500      | 3.27s  | 0.4166         | 0.3308        | 0.3139          | 0.5000         |
| 1000     | 7.09s  | 0.4170         | 0.3320        | 0.3088          | 0.5000         |
| 2000     | 13.91s | 0.4179         | 0.3324        | 0.3019          | 0.5000         |

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
