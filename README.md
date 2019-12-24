# observation-controller

Sample controller to create Observations based on configuration.

## What's an Observation, anyway

See
[this design document](https://docs.google.com/document/d/1TvHt0QKAKcvrqJTQVjNfFuVa5hVZQhCzTn7QDltANKk/edit#heading=h.r8nd4wsy59d5)
for more details.

In short, an Observation is a Kubernetes object which represents information
about a Kubernetes resource which is tracked in an external system. The primary
intent is to make it easy to navigate to other views of the same information
which don't necessarily make sense to descibe as Kubernetes API resources (for
example, container logs or application dashboards).
