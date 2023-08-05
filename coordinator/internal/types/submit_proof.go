package types

// SubmitProofParameter the SubmitProof api request parameter
type SubmitProofParameter struct {
	TaskID   string `form:"task_id" json:"task_id" binding:"required"`
	TaskType int    `form:"task_type" json:"task_type" binding:"required"`
	Status   int    `form:"status" json:"status"`
	Proof    string `form:"proof" json:"proof"`
}