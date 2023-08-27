package clients





type Clients struct {
	UserClient
}

func NewApiClients(uc UserClient) *Clients {
	return &Clients{
		UserClient: uc,
	}
}