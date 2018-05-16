# Dimension Recombination Reporter - Controller

WORK IN PROGRESS - 10% project.

The DRR is a microservice based project to analyse all permutations from structural alterations (combinations of dimension items added/removed) that could be applied to a given source dataset. Presenting the findings as a human readable report.

The result writer checks each result, discarding those that don't meet the filter (for now hardcoded as those that don't reduce sparsity) and writing positive results to CSV before sending back to the controller for further permutations.

## Controller

The controller takes results from the results queue. If the result is positive (i.e reduces sparsity) it sends that result to the result writer and generates further premutations to try which are sent to the task queue.

It also tracks results recieved vs results expected and send a "job competed" message when the two match.

It discards non-positive results.
