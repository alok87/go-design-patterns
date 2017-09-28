package k8s

// Factory provides abstractions that allow the Kubectl command to be extended across multiple types
// of resources and different API sets.
type Factory interface {
    ClientAccessFactory
    ObjectMappingFactory
    BuilderFactory
}

// ClientAccessFactory holds the first level of factory methods.
// Generally provides discovery, negotiation, and no-dep calls.
type ClientAccessFactory interface {

}

// ObjectMappingFactory holds the second level of factory methods.  These functions depend upon ClientAccessFactory methods.
// Generally they provide object typing and functions that build requests based on the negotiated clients.
type ObjectMappingFactory interface {

}

// BuilderFactory holds the second level of factory methods.  These functions depend upon ObjectMappingFactory and ClientAccessFactory methods.
// Generally they depend upon client mapper functions
type BuilderFactory interface {

}
