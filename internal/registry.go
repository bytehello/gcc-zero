package internal

type Registry struct {
}

func (r *Registry) GetConn() (EtcdClient, error) {
	return nil, nil
}

type cluster struct {
}
