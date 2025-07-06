package factory

type Model struct {
	// Id          string `json:"id"`          //mongodb default _id
	Request_id  string `json:"request_id"`  //request id
	Author      string `json:"author"`      //author of the model
	Action      string `json:"action"`      //action to be performed
	From_branch string `json:"from_branch"` //branch from which the model is created
	To_branch   string `json:"to_branch"`   //branch to which the model is created
	Created_at  string `json:"created_at"`  //created at
}
