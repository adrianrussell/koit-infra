# Koit-Infra

Project to learn go by writing program that deploys the same infrastructure to both vultr and aws.

Infrastructure to deploy

- Virtual Compute Instance
- Load Balancer

Considerations

- api design - looking to provide a fluent api for use to configure infrastructure across providers
- secret management - need to externalize the management of secrets
- testing - provide unit testing for interfaces and see what integration testing can implemented to verify deployed configuration.
