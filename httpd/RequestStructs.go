package httpd

type S3InitRequest struct {
	Filename string `json:"filename"`
	Multi    bool   `json:"multi"`
}
type S3InitResponse struct {
	Passkey string `json:"passkey"`
}

type URLResponse struct {
	URL string `json:"url"`
}

type Template struct {
	Yaml string `json:"yaml"`
	Name string `json:"name"`
}
type ApplyReturn struct {
	Status  string `json:"status"`
	Missing string `json:"missing,omitempty"`
	WfName  string `json:"wfname,omitempty"`
}

type ArtifactResponse struct {
	Pod string `json:"pod"`
	URL string `json:"url"`
}

type WFStatus struct {
	Workflow      string `json:"workflow"`
	Status        string `json:"status"`
	Statusmessage string `json:"statusmessage"`
	Running       int    `json:"steps"`
	Finished      int    `json:"finished"`
}

type SchemeInfo struct {
	Name       string            `json:"name"`
	Yaml       string            `json:"yaml"`
	Parameters map[string]string `json:"parameters"`
}
