# hairpin-test

Continuously performs a http request to an endpoint over a period of 10 seconds with the intent 
of discovering hairpin issues.  This is accomplished by deploying to the test to each node which 
is a target in the load balancer.   

## Building

podman build . --tag <your registry>/hairpin-test:latest

## Running

1. Update `job.yaml` with the load balancer you wish to test.
2. Update the `nodeSelector` ensure the test only runs on the nodes of interest.  By default,
this test runs on all three control plane nodes.
3. Update the parallelism to ensure the correct number of pods are deployed.
4. Apply `job.yaml`.
5. Wait for the job to complete.
