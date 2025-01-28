# VidFlex

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

A Q-Learning-based adaptive video streaming optimizer, entirely written in Go, designed to maximize user satisfaction while minimizing bandwidth costs. The current implementation operates on synthetic data with simplified states and utilizes Q-Learning with an epsilon-greedy policy.

## Benchmark Results

**Device**: 15-inch MacBook Air with M3 Chip (8-Core CPU, 10-Core GPU, and 16GB Unified Memory)

| Episodes | Users | Time (s) | Dynamic Reward | Static Reward | Dynamic Penalty | Static Penalty |
| -------- | ----- | -------- | -------------- | ------------- | --------------- | -------------- |
| 500      | 500   | 1.09     | 0.3447         | 0.3420        | 0.4348          | 0.4284         |
| 500      | 1000  | 2.19     | 0.3457         | 0.3420        | 0.4363          | 0.4283         |
| 500      | 2000  | 4.75     | 0.3459         | 0.3420        | 0.4403          | 0.4488         |
| 1000     | 500   | 2.17     | 0.3438         | 0.3420        | 0.4432          | 0.4562         |
| 1000     | 1000  | 4.59     | 0.3441         | 0.3420        | 0.4357          | 0.4357         |
| 1000     | 2000  | 9.57     | 0.3445         | 0.3420        | 0.4375          | 0.4404         |
| 2000     | 500   | 4.46     | 0.3428         | 0.3420        | 0.4614          | 0.4430         |
| 2000     | 1000  | 9.27     | 0.3428         | 0.3420        | 0.4405          | 0.4406         |
| 2000     | 2000  | 19.68    | 0.3432         | 0.3420        | 0.4333          | 0.4403         |

## Features

- [x] Simulate concurrent users watching the same video stream with randomized network bandwidth ranges (0–1 Mbps, 1–3 Mbps, 3–5 Mbps, 5+ Mbps) and device types (mobile, tablet, desktop/TV).
- [x] Adjust video quality by modifying one metric at a time (resolution, bitrate, frame rate) through discrete actions.
- [x] Update the policy using a reward function that combines engagement score (prioritizing higher quality), bandwidth cost penalties, and a collaborative adjustment based on the average satisfaction of other concurrent users in similar states.  
       Formula:  
       $$Q(s, a) \leftarrow Q(s, a) + \alpha \left[\text{reward} + \gamma \max_{a'} Q(s', a') + 0.2 \times \text{average satisfaction}\right]$$

- [x] Implement an epsilon-greedy policy with an initial epsilon = 0.8 and epsilon decay of 5% per 100 episodes.
- [x] **System assumptions**: Concurrent users are managed by a single agent via `goroutines`; simulated bandwidth dips reduce the bandwidth tier by one level with 10% probability.
- [x] Train the agent and benchmark performance against a static policy (fixed 720p, 1 Mbps, 30 fps) via command line (see the below section for more details).

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

3. To generate benchmark results, use the precreated script files:

   ```bash
   chmod +x benchmark.sh
   ./benchmark.sh
   ```

   Check the results.csv file in the root directory for the benchmark results.

_Note: The static policy (fixed 720p/1 Mbps/30 fps) is automatically included in the benchmark for comparison._

## Future Roadmap

- [ ] Implement advanced algorithms like Deep Q-Networks (DQN) to stabilize learning as user numbers scale, using techniques like experience replay and target networks.
- [ ] Integrate alternative exploration strategies (e.g., Upper Confidence Bound [UCB], Thompson Sampling) to improve state-action space exploration beyond epsilon-greedy.
- [ ] Expand state definitions with real-world network parameters (e.g., latency, jitter, packet loss) and device-specific capabilities (e.g., codec support, screen resolution).

## License

This project is licensed under the GNU General Public License v3.0 . See the [`LICENSE`](/LICENSE) file for details.
