package strategy

type EndpointStrategy struct {
	EventsEndpointStrategy
	DomainsEndpointStrategy
}

func NewEndpointStrategy(versionString string) EndpointStrategy {
	version, err := ParseVersion(versionString)
	if err != nil {
		version = Version{0, 0, 0}
	}

	strategy := EndpointStrategy{}

	if version.GreaterThanOrEqualTo(Version{2, 2, 0}) {
		strategy.EventsEndpointStrategy = globalEventsEndpointStrategy{}
		strategy.DomainsEndpointStrategy = orgScopedDomainsEndpointStrategy{}
	} else {
		strategy.EventsEndpointStrategy = appScopedEventsEndpointStrategy{}
		strategy.DomainsEndpointStrategy = globalDomainsEndpointStrategy{}
	}

	return strategy
}