package sls_store

type slsTraceInstance interface {
	traceLogStore() string
	serviceDependencyLogStore() string
}

func newSlsTraceInstance(instance string) slsTraceInstance {
	return &slsTraceInstanceImpl{
		instance:                      instance,
		traceLogStoreName:             instance + "-traces",
		serviceDependencyLogStoreName: instance + "-traces-deps",
	}
}

type slsTraceInstanceImpl struct {
	instance                      string
	traceLogStoreName             string
	serviceDependencyLogStoreName string
}

func (s *slsTraceInstanceImpl) traceLogStore() string {
	return s.traceLogStoreName
}

func (s *slsTraceInstanceImpl) serviceDependencyLogStore() string {
	return s.serviceDependencyLogStoreName
}
