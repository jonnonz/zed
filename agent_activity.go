package zd

type AgentsActivity struct {
	AgentId         *int    `json:"agent_id"`
	Status          *string `json:"status"`
	AvailableTime   *int    `json:"available_time"`
	CallsAccepted   *int    `json:"calls_accepted"`
	CallsDenied     *int    `json:"calls_denied"`
	CallsMissed     *int    `json:"calls_missed"`
	AverageTalkTime *int    `json:"average_talk_time"`
}

type ActivityResponse struct {
	Activities *[]AgentsActivity `json:"users"`
	Next       *string           `json:"next_page,omitempty"`
	Previous   *string           `json:"previous_page,omitempty"`
	Count      *int              `json:"count,omitempty"`
}

type ActivityService struct {
	client *Client
}

func (s *ActivityService) GetActivity() ([]AgentsActivity, error) {
	resource := make([]AgentsActivity, 0)

	rp, next, _, err := s.getPage("")
	if err != nil {
		return nil, err
	}

	resource = append(resource, *rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return nil, err
		}
		next = nx
		resource = append(resource, *rp...)
	}

	return resource, err
}

func (s *ActivityService) getPage(url string) (*[]AgentsActivity, *string, *Response, error) {

	if url == "" {
		url = "users.json?role=agent"
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(ActivityResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	next := response.Next
	resource := response.Activities
	return resource, next, resp, err
}