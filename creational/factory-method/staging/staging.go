package staging

type Staging interface {
    Create() bool
    Delete() bool
    Update() bool
}

type Product struct {
    productoptions ProductOptions
}

type ProductOptions struct {
    name string
    repo string
    admins []string
    testedCommit string
}

type SandBoxEnvOptions struct {
    name string
    products []Product
}

type LoadTestEnvOptions struct {
    name string
    products []Product
}

type SandboxEnv struct {
    sandBoxEnvOptions SandBoxEnvOptions
}

type LoadTestEnv struct {
    loadTestEnvOptions LoadTestEnvOptions
}

# Factory method
func NewProduct(productoptions ProductOptions) {
    return &Product{
        productoptions: ProductOptions
    }
}

# Factory method
func NewSandboxEnv(options NewSandBoxEnvOptions) {
    return &SandboxEnv{
    }
}
