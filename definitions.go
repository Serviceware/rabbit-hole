package rabbithole

// ExportedDefinitions represents definitions exported from a RabbitMQ cluster
type ExportedDefinitions struct {
	RabbitVersion    string `json:"rabbit_version"`
	RabbitMQVersion  string `json:"rabbitmq_version"`
	ProductName      string `json:"product_name"`
	ProductVersion   string `json:"product_version"`
	Users            []UserInfo
	Vhosts           []VhostInfo
	Permissions      []Permissions
	TopicPermissions []TopicPermissionInfo
	Parameters       []RuntimeParameter
	GlobalParameters interface{} `json:"global_parameters"`
	Policies         []PolicyDefinition
	Queues           []QueueInfo
	Exchanges        []ExchangeInfo
	Bindings         []BindingInfo
}

//
// GET /api/definitions
//

func (c *Client) ExportDefinitions() (p *ExportedDefinitions, err error) {
	req, err := newGETRequest(c, "definitions")
	if err != nil {
		return nil, err
	}

	if err = executeAndParseRequest(c, req, &p); err != nil {
		return nil, err
	}

	return p, nil
}
