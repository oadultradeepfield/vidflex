#!/bin/bash

# Navigate to the cmd/train directory
cd cmd/train

# Build the application
go build -o train

# Create results file header
echo "Episodes,Users,Time,DynamicReward,StaticReward,DynamicPenalty,StaticPenalty" > ../../results.csv

# Run benchmarks with different configurations
for episodes in 500 1000 2000; do
    for users in 500 1000 2000; do
        echo "Running benchmark with $episodes episodes and $users users..."
        
        # Run the program and capture output
        output=$(./train -episodes=$episodes -users=$users 2>&1)
        
        # Extract metrics using the new FINAL_METRICS line
        metrics=$(echo "$output" | grep "FINAL_METRICS")
        
        if [ -z "$metrics" ]; then
            echo "Error: No metrics found in output for $episodes episodes and $users users."
            continue
        fi

        # Parse metrics
        time=$(echo "$output" | grep "Training completed" | awk '{print $4}')
        dynamic_reward=$(echo "$metrics" | awk -F'|' '{print $4}' | awk -F':' '{print $2}')
        static_reward=$(echo "$metrics" | awk -F'|' '{print $5}' | awk -F':' '{print $2}')
        dynamic_penalty=$(echo "$metrics" | awk -F'|' '{print $6}' | awk -F':' '{print $2}')
        static_penalty=$(echo "$metrics" | awk -F'|' '{print $7}' | awk -F':' '{print $2}')

        # Append to CSV (saved in the root directory)
        echo "$episodes,$users,$time,$dynamic_reward,$static_reward,$dynamic_penalty,$static_penalty" >> ../../results.csv
    done
done

# Return to the original directory
cd ../..

echo "Benchmark completed. Results saved to results.csv"